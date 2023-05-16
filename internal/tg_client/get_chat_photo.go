package tg_client

import (
	"bytes"
	"context"
	"fmt"
	"syscall"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgerr"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	pkgErrors "github.com/pkg/errors"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tgInfo) GetChatPhotoFromApi(filename *string, chat Chat) (string, error) {
	photoLink, err := t.getChatPhotoFlow(filename, chat)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTgAsInterface(t.cfg, t.ctx).(TelegramClient)
			t.superUserClient = cl.GetTg().superUserClient
			return t.GetChatPhotoFromApi(filename, chat)
		}

		duration, isFlood := tgerr.AsFloodWait(err)
		if isFlood {
			t.log.Warnf("we need to wait `%s`", duration)
			time.Sleep(duration)
			return t.GetChatPhotoFromApi(filename, chat)
		}

		return "", errors.Wrap(err, fmt.Sprintf("failed to get chat Photo"))
	}

	t.log.Infof("successfully got chat Photo")
	return photoLink, nil
}

func (t *tgInfo) getChatPhotoFlow(filename *string, chat Chat) (string, error) {
	if filename != nil {
		return t.handlePhoto(*filename, chat)
	}

	return t.handlePhoto(uuid.New().String(), chat)
}

func (t *tgInfo) handlePhoto(filename string, chat Chat) (string, error) {
	if chat.Photo == nil {
		return "", errors.New("no photo to handle")
	}

	inputPeer := tg.InputPeerClass(nil)
	if chat.AccessHash != nil {
		inputPeer = &tg.InputPeerChannel{ChannelID: chat.Id, AccessHash: *chat.AccessHash}
	} else {
		inputPeer = &tg.InputPeerChat{ChatID: chat.Id}
	}

	switch converted := (*chat.Photo).(type) {
	case *tg.ChatPhotoEmpty:
		t.log.Warnf("no Photo set for `%s`", chat.Title)
		return "", nil
	case *tg.ChatPhoto:
		file, err := t.downloadPhoto(converted, inputPeer)
		if err != nil {
			return "", err
		}

		return t.processUploadFile(file, filename)
	default:
		return "", errors.New("unknown photo type")
	}
}

func (t *tgInfo) processUploadFile(fileClass *tg.UploadFileClass, filename string) (string, error) {
	if fileClass == nil {
		return "", errors.New("fileClass is empty")
	}

	switch file := (*fileClass).(type) {
	case *tg.UploadFile:
		fileType, err := t.getFileType(file.Bytes)
		if err != nil {
			return "", err
		}
		return storeRemotePhoto(file.Bytes, filename+fileType, t.cfg.Aws())
	case *tg.UploadFileCDNRedirect:
		return t.handleCDNRedirect(file, filename)
	default:
		return "", errors.New("unknown upload file type")
	}
}

func (t *tgInfo) handleCDNRedirect(cdnRedirectFile *tg.UploadFileCDNRedirect, filename string) (string, error) {
	cdnFile, err := t.downloadCDNPhoto(cdnRedirectFile.FileToken)
	if err != nil {
		return "", err
	}

	if cdnFile == nil {
		return "", errors.New("cdn file is empty")
	}

	switch file := (*cdnFile).(type) {
	case *tg.UploadCDNFile:
		fileType, err := t.getFileType(file.Bytes)
		if err != nil {
			return "", err
		}

		return storeRemotePhoto(file.Bytes, filename+fileType, t.cfg.Aws())
	default:
		return "", errors.New("failed to handle cdn redirect file")
	}
}

func (t *tgInfo) downloadPhoto(photo *tg.ChatPhoto, peer tg.InputPeerClass) (*tg.UploadFileClass, error) {
	file, err := t.superUserClient.API().UploadGetFile(t.ctx, &tg.UploadGetFileRequest{
		Precise:      true,
		CDNSupported: false,
		Location: &tg.InputPeerPhotoFileLocation{
			PhotoID: photo.PhotoID,
			Flags:   photo.Flags,
			Peer:    peer,
			Big:     false,
		},
		Limit:  1024 * 1024,
		Offset: 0,
	})
	if err != nil {
		return nil, err
	}

	return &file, err
}

func (t *tgInfo) downloadCDNPhoto(token []byte) (*tg.UploadCDNFileClass, error) {
	file, err := t.superUserClient.API().UploadGetCDNFile(t.ctx, &tg.UploadGetCDNFileRequest{
		FileToken: token,
		Offset:    0,
		Limit:     1024 * 1024,
	})
	if err != nil {
		return nil, err
	}

	return &file, nil
}

func (t *tgInfo) getFileType(fileBytes []byte) (string, error) {
	mType := mimetype.Detect(fileBytes)
	switch mimetype.Detect(fileBytes).String() {
	case "image/jpeg":
		return ".jpeg", nil
	case "image/png":
		return ".png", nil
	default:
		return "", errors.Errorf("unknown image type `%s`", mType.String())
	}
}

func storeRemotePhoto(imgBytes []byte, filename string, awsCfg *config.AwsCfg) (string, error) {
	s3Client, err := minio.New("s3.amazonaws.com", &minio.Options{
		Creds:  credentials.NewStaticV4(awsCfg.Id, awsCfg.Secret, ""),
		Secure: true,
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to create new aws s3 client")
	}

	_, err = s3Client.PutObject(
		context.Background(),
		data.S3BucketName,
		"telegram-module/"+filename,
		bytes.NewReader(imgBytes),
		-1,
		minio.PutObjectOptions{},
	)
	if err != nil {
		return "", errors.Wrap(err, "failed to put file in bucket")
	}

	return data.S3BucketEndpoint + "/telegram-module/" + filename, nil
}
