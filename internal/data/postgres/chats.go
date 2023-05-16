package postgres

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	chatsTableName        = "chats"
	chatsTitleColumn      = chatsTableName + ".title"
	chatsIdColumn         = chatsTableName + ".id"
	chatsAccessHashColumn = chatsTableName + ".access_hash"
)

type ChatsQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewChatsQ(db *pgdb.DB) data.Chats {
	return &ChatsQ{
		db:            db,
		selectBuilder: sq.Select(chatsTableName + ".*").From(chatsTableName),
		deleteBuilder: sq.Delete(chatsTableName),
	}
}

func (r ChatsQ) New() data.Chats {
	return NewChatsQ(r.db)
}

func (r ChatsQ) Get() (*data.Chat, error) {
	var result data.Chat
	err := r.db.Get(&result, r.selectBuilder)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (r ChatsQ) Select() ([]data.Chat, error) {
	var result []data.Chat

	err := r.db.Select(&result, r.selectBuilder)

	return result, err
}

func (r ChatsQ) Upsert(chat data.Chat) error {
	updateStmt, args := sq.Update(" ").
		Set("title", chat.Title).
		Set("members_amount", chat.MembersAmount).
		Set("photo_link", chat.PhotoLink).
		Set("photo_name", chat.PhotoName).
		MustSql()

	query := sq.Insert(chatsTableName).SetMap(structs.Map(chat)).
		Suffix("ON CONFLICT (id, access_hash) DO "+updateStmt, args...)

	err := r.db.Exec(query)
	return errors.Wrap(err, "failed to insert chat")
}

func (r ChatsQ) Delete() error {
	var deleted []data.Chat

	err := r.db.Select(&deleted, r.deleteBuilder.Suffix("RETURNING *"))
	if err != nil {
		return err
	}

	if len(deleted) == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r ChatsQ) FilterByTitles(titles ...string) data.Chats {
	equalTitles := sq.Eq{chatsTitleColumn: titles}
	r.selectBuilder = r.selectBuilder.Where(equalTitles)
	r.deleteBuilder = r.deleteBuilder.Where(equalTitles)

	return r
}

func (r ChatsQ) FilterByIds(ids ...int64) data.Chats {
	equalIds := sq.Eq{chatsIdColumn: ids}
	r.selectBuilder = r.selectBuilder.Where(equalIds)
	r.deleteBuilder = r.deleteBuilder.Where(equalIds)

	return r
}

func (r ChatsQ) FilterByAccessHash(accessHash *int64) data.Chats {
	equalAccessHash := sq.Eq{chatsAccessHashColumn: accessHash}
	r.selectBuilder = r.selectBuilder.Where(equalAccessHash)
	r.deleteBuilder = r.deleteBuilder.Where(equalAccessHash)

	return r
}
