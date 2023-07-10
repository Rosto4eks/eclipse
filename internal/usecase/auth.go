package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/Rosto4eks/eclipse/internal/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
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

func (u *usecase) GetUserByName(name string) (models.User, error) {
	return u.database.GetUserByName(name)
}

func (u *usecase) SignIn(name, password string) error {
	usr, err := u.database.GetUserByName(name)
	if err != nil {
		return err
	} else if usr.Password == hash(password) {
		return nil
	}
	return errors.New("there are no user with such credits")
}

func (u *usecase) GenerateToken(name, password, role string) (string, error) {
	err := u.SignIn(name, password)
	if err != nil {
		return "", err
	}

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

func (u *usecase) ParseToken(stringToken string, signingKey []byte) (string, error) {
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
		//if time.Now().Unix() > claims.ExpiresAt {
		//
		//}
		return claims.Username, nil
	}
	return "", errors.New("invalid access token")
}

func (u *usecase) WriteCookie(token string, ctx echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "jwt_token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	ctx.SetCookie(cookie)
	return ctx.String(http.StatusOK, "write a cookie")
}

func (u *usecase) ReadCookie(ctx echo.Context) (string, error) {
	cookie, err := ctx.Cookie("jwt_token")
	if err != nil {
		return "", err
	}
	return cookie.Value, ctx.String(http.StatusOK, "read a cookie")
}

func hash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
