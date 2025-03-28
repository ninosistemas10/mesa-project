package login

import (
	"database/sql"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/ninosistemas10/delivery/domain/login"
	"github.com/ninosistemas10/delivery/infrastructure/handler/response"
	"github.com/ninosistemas10/delivery/model"
)

type handler struct {
	useCase   login.UseCase
	responser response.API
}

func newHandler(useCase login.UseCase) handler {
	return handler{useCase: useCase}
}

func (h handler) Login(c echo.Context) error {
	m := model.Login{}
	err := c.Bind(&m)
	if err != nil {
		return h.responser.BindFailed(err)
	}

	u, t, err := h.useCase.Login(m.Email, m.Password, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		if strings.Contains(err.Error(), "bcrypt.CompareHashAndPassword()") ||
			errors.Is(err, sql.ErrNoRows) {
			resp := model.MessageResponse{
				Data:     "wrong user or password",
				Messages: model.Responses{{Code: response.AuthError, Message: "wrong user or password"}},
			}
			return c.JSON(http.StatusBadRequest, resp)
		}
		return h.responser.Error(c, "useCase.Login()", err)
	}

	return c.JSON(h.responser.OK(map[string]interface{}{"user": u, "token": t}))
}


