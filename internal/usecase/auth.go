package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/Rosto4eks/eclipse/internal/models"
	"github.com/golang-jwt/jwt"
)

const signingKey = "sw4567tyuhgftghiEWGwhufI&#$"

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func (u *usecase) SignUp(usr models.User) (string, error) {
	usr.Password = hash(usr.Password)
	if err := u.database.AddUser(usr); err != nil {
		return "", err
	}
	return generateToken(usr.Name, usr.Role)
}

func (u *usecase) GetUserByName(name string) (models.User, error) {
	return u.database.GetUserByName(name)
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

func (u *usecase) Auth(token, role string) error {
	usr, err := parseToken(token)
	if err != nil {
		return err
	}
	if usr.Role != role {
		return errors.New("permission denied")
	}
	return nil
}

func (u *usecase) AuthHeader(token string) string {
	usr, err := parseToken(token)
	if err != nil {
		return "sign in"
	}
	return usr.Name
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}

func parseToken(stringToken string) (models.User, error) {
	token, err := jwt.ParseWithClaims(stringToken, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}
			return []byte(signingKey), nil
		})
	if err != nil {
		return models.User{}, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		usr := models.User{
			Name: claims.Username,
			Role: claims.Role,
		}
		return usr, nil
	}
	return models.User{}, errors.New("invalid access token")
}

func hash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
