package helpers

import (
	"github.com/golang-jwt/jwt"
	"gitlab.com/distributed_lab/acs/auth/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func GenerateAccessToken(user data.User, expires int64, secret string, permissions string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expires
	claims["owner_id"] = user.Id
	claims["email"] = user.Email
	claims["module.permission"] = permissions

	return token.SignedString([]byte(secret))
}

func GenerateRefreshToken(user data.User, expires int64, secret string, permissions string) (string, error, jwt.MapClaims) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expires
	claims["owner_id"] = user.Id
	claims["email"] = user.Email
	claims["module.permission"] = permissions

	signedToken, err := token.SignedString([]byte(secret))

	return signedToken, err, claims
}

func CheckRefreshToken(tokenStr string, ownerId int64, secret string) error {
	token, err := parse(tokenStr, []byte(secret))
	if err != nil {
		return errors.Wrap(err, "some error while parsing jwt token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if int64(claims["owner_id"].(float64)) != ownerId {
			return errors.New("invalid token")
		}
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return err
}

func parse(tokenStr string, secret []byte) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return secret, nil
	})
}

func CheckValidToken(tokenStr string, secret string) error {
	token, err := parse(tokenStr, []byte(secret))
	if err != nil {
		return errors.Wrap(err, "some error while parsing jwt token")
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return err
}

func ParseJwtToken(tokenStr string, secret string) (jwt.MapClaims, error) {
	token, err := parse(tokenStr, []byte(secret))
	if err != nil {
		return nil, errors.Wrap(err, "some error while parsing jwt token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
