package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/Rosto4eks/eclipse/internal/models"
	"github.com/labstack/echo/v4"
)

func (h *handler) GetSignIn(ctx echo.Context) error {
	headerName := h.authHeader(ctx)
	return ctx.Render(200, "auth.html", map[string]interface{}{
		"header": headerName,
		"type":   "signin",
	})
}

func (h *handler) GetSignUp(ctx echo.Context) error {
	headerName := h.authHeader(ctx)
	return ctx.Render(200, "auth.html", map[string]interface{}{
		"header": headerName,
		"type":   "signup",
	})
}

func (h *handler) GetLogOut(ctx echo.Context) error {
	h.cleanJWT(ctx)
	return ctx.Redirect(302, "/")
}

func (h *handler) PostSignUp(ctx echo.Context) error {
	headerName := h.authHeader(ctx)

	usr := models.User{
		Name:     ctx.FormValue("name"),
		Password: ctx.FormValue("password"),
	}
	token, err := h.usecase.SignUp(usr)
	if err != nil {
		return ctx.Render(401, "auth.html", map[string]interface{}{
			"header": headerName,
			"type":   "signup",
			"error":  err.Error(),
		})
	}
	h.writeJWT(token, ctx)
	return ctx.Redirect(302, "/?welcome=true")
}

func (h *handler) PostSignIn(ctx echo.Context) error {
	headerName := h.authHeader(ctx)

	Name := ctx.FormValue("name")
	Password := ctx.FormValue("password")

	token, err := h.usecase.SignIn(Name, Password)
	if err != nil {
		return ctx.Render(401, "auth.html", map[string]interface{}{
			"header": headerName,
			"type":   "signin",
			"error":  err.Error(),
		})
	}
	h.writeJWT(token, ctx)
	return ctx.Redirect(302, "/?welcome=true")
}

func (h *handler) auth(ctx echo.Context, role string) error {
	token, err := h.readJWT(ctx)
	if err != nil {
		return err
	}

	return h.usecase.Auth(token, role)
}

func (h *handler) authHeader(ctx echo.Context) string {
	token, err := h.readJWT(ctx)
	if err != nil {
		return "sign in"
	}
	return h.usecase.AuthHeader(token)
}

func (h *handler) writeJWT(token string, ctx echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "jwt_token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	ctx.SetCookie(cookie)
}

func (h *handler) readJWT(ctx echo.Context) (string, error) {
	cookie, err := ctx.Cookie("jwt_token")
	if cookie == nil {
		return "", errors.New("cookie not found")
	}
	return cookie.Value, err
}

func (h *handler) cleanJWT(ctx echo.Context) error {
	cookie, err := ctx.Cookie("jwt_token")
	if err != nil {
		return err
	}
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	ctx.SetCookie(cookie)
	return nil
}
