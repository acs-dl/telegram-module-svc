package helpers

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"gitlab.com/distributed_lab/acs/auth/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func GenerateAccessToken(dataToGenerate data.GenerateTokens) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	_, err := setMapClaimsFromStructure(dataToGenerate, token.Claims.(jwt.MapClaims))
	if err != nil {
		return "", errors.New("failed to set claims")
	}

	return token.SignedString([]byte(dataToGenerate.Secret))
}

func GenerateRefreshToken(dataToGenerate data.GenerateTokens) (string, error, *data.JwtClaims) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claimsStruct, err := setMapClaimsFromStructure(dataToGenerate, claims)
	if err != nil {
		return "", errors.New("failed to set claims"), nil
	}

	signedToken, err := token.SignedString([]byte(dataToGenerate.Secret))
	if err != nil {
		return "", errors.Wrap(err, "failed to get signed string"), nil
	}

	return signedToken, err, claimsStruct
}

func CheckValidityAndOwnerForRefreshToken(tokenStr string, ownerId int64, secret string) error {
	token, err := parse(tokenStr, []byte(secret))
	if err != nil {
		return errors.Wrap(err, "some error while parsing jwt token")
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid token")
	}

	claimsStruct, err := getClaimsStructureFromMap(claims)
	if err != nil {
		return errors.Wrap(err, "failed to get claims structure from map")
	}

	if claimsStruct.OwnerId != ownerId {
		return errors.New("invalid token")
	}

	return nil
}

func parse(tokenStr string, secret []byte) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return secret, nil
	})
}

func CheckTokenValidity(tokenStr string, secret string) (*jwt.Token, error) {
	token, err := parse(tokenStr, []byte(secret))
	if err != nil {
		return nil, errors.Wrap(err, "some error while parsing jwt token")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

func RetrieveClaimsFromJwtString(tokenStr string, secret string) (*data.JwtClaims, error) {
	token, err := CheckTokenValidity(tokenStr, secret)
	if err != nil {
		return nil, errors.New("failed to check token validity")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	claimsStruct, err := getClaimsStructureFromMap(claims)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get claims structure from map")
	}

	return claimsStruct, err
}

func getClaimsStructureFromMap(claims jwt.MapClaims) (*data.JwtClaims, error) {
	jsonClaims, err := json.Marshal(claims)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal claims")
	}

	var claimsStruct data.JwtClaims
	err = json.Unmarshal(jsonClaims, &claimsStruct)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal claims")
	}

	return &claimsStruct, nil
}

func setMapClaimsFromStructure(dataToGenerate data.GenerateTokens, claims jwt.MapClaims) (*data.JwtClaims, error) {
	claimsStruct := data.JwtClaims{
		ExpiresAt:        dataToGenerate.AccessLife,
		OwnerId:          dataToGenerate.User.Id,
		Email:            dataToGenerate.User.Email,
		ModulePermission: dataToGenerate.PermissionsString,
	}

	jsonClaims, err := json.Marshal(claimsStruct)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal claims")
	}

	err = json.Unmarshal(jsonClaims, &claims)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal claims")
	}

	return &claimsStruct, nil
}
