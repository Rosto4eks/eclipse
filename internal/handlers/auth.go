package handlers

import (
	"net/http"
	"time"

	"github.com/Rosto4eks/eclipse/internal/models"
	"github.com/labstack/echo/v4"
)

func (h *handler) GetSignIn(ctx echo.Context) error {
	return ctx.Render(200, "auth.html", map[string]interface{}{"type": "signin"})
}

func (h *handler) GetSignUp(ctx echo.Context) error {
	return ctx.Render(200, "auth.html", map[string]interface{}{"type": "signup"})
}

func (h *handler) PostSignUp(ctx echo.Context) error {
	usr := models.User{
		Name:     ctx.FormValue("name"),
		Password: ctx.FormValue("password"),
	}
	token, err := h.usecase.SignUp(usr)
	if err != nil {
		return ctx.Render(401, "auth.html", map[string]interface{}{
			"type":  "signup",
			"error": err.Error(),
		})
	}
	h.writeJWT(token, ctx)
	return ctx.Redirect(301, "/")
}

func (h *handler) PostSignIn(ctx echo.Context) error {
	Name := ctx.FormValue("name")
	Password := ctx.FormValue("password")

	token, err := h.usecase.SignIn(Name, Password)
	if err != nil {
		return ctx.Render(401, "auth.html", map[string]interface{}{
			"type":  "signin",
			"error": err.Error(),
		})
	}
	h.writeJWT(token, ctx)
	return ctx.Redirect(301, "/")
}

func (h *handler) auth(ctx echo.Context, role string) error {
	token, err := h.readJWT(ctx)
	if err != nil {
		return err
	}

	return h.usecase.Auth(token, role)
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
	return cookie.Value, err
}
