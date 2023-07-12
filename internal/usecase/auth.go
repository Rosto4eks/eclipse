package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/Rosto4eks/eclipse/internal/models"
	"github.com/golang-jwt/jwt"
	"time"
)

const SigningKey = "sw4567tyuhgftghiEWGwhufI&#$"

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func (u *usecase) NewUser(usr models.User) error {
	usr.Password = hash(usr.Password)
	return u.database.AddUser(usr)
}

func (u *usecase) SignIn(name, password string) (string, error) {
	usr, err := u.database.GetUserByName(name)
	if err != nil {
		return "", err
	} else if usr.Password == hash(password) {
		token, err := generateToken(name, usr.Role)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return "", errors.New("there are no user with such credits")
}

func generateToken(name, role string) (string, error) {
	claims := Claims{
		name,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(SigningKey))
}

func parseToken(stringToken string, signingKey []byte) (string, error) {
	token, err := jwt.ParseWithClaims(stringToken, Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}
			return signingKey, nil
		})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Username, nil
	}
	return "", errors.New("invalid access token")
}

func hash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
