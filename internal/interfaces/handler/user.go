package handler

import (
	"net/http"

	"golang-webserver-practise/internal/domain/model"
	"golang-webserver-practise/internal/registory"
	userUsecase "golang-webserver-practise/internal/usecase/user"
	"golang-webserver-practise/pkg/converter"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Index(c echo.Context) error
	Show(c echo.Context) error
	Create(c echo.Context) error
}

type userImpl struct {
	repo registory.Repository
}

type ValidationErrorResponse struct {
	Key   string
	Error string
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

func (impl *userImpl) Create(c echo.Context) error {
	u := new(model.User)

	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, &ValidationErrorResponse{Key: "User.Email", Error: err.Error()})
	}

	uc := userUsecase.NewCreateImpl(impl.repo, u)
	flg, err := uc.IsUnique()
	if err != nil {
		return err
	}
	if !flg {
		return c.JSON(http.StatusBadRequest, &ValidationErrorResponse{Key: "User.Email", Error: "not unique."})
	}

	u, err = uc.Exec()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
