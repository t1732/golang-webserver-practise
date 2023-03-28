package handler

import (
	"net/http"

	"golang-webserver-practise/internal/registory"
	"golang-webserver-practise/pkg/converter"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Index(c echo.Context) error
	Show(c echo.Context) error
}

type userImpl struct {
	repo registory.Repository
}

func NewUserImpl(repo registory.Repository) UserHandler {
	return &userImpl{repo}
}

func (impl *userImpl) Index(c echo.Context) error {
	users, err := impl.repo.NewUserRepository().All()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (impl *userImpl) Show(c echo.Context) error {
	user, err := impl.repo.NewUserRepository().FindById(converter.StringToInt64(c.Param("id")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
