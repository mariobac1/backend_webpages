package user

import (
	"errors"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/mariobac1/backend_webpages/domain/user"
	"github.com/mariobac1/backend_webpages/infrastructure/handler/response"
	"github.com/mariobac1/backend_webpages/model"
)

type handler struct {
	useCase   user.UseCase
	responser response.API
}

func newHandler(uc user.UseCase) handler {
	return handler{useCase: uc}

}

func (h handler) Create(c echo.Context) error {
	m := model.User{}

	if err := c.Bind(&m); err != nil {
		return h.responser.BindFailed(err)
	}

	if err := h.useCase.Create(&m); err != nil {
		return h.responser.Error(c, "useCase.Create()", err)
	}
	return c.JSON(h.responser.Created(m))
}

// MySelf returns the data from my profile
func (h handler) MySelf(c echo.Context) error {
	ID, ok := c.Get("userID").(uuid.UUID)
	if !ok {
		return h.responser.Error(c, "c.Get().(uuid.UUID)", errors.New("couldn´t parse the ID"))
	}

	u, err := h.useCase.GetByID(ID)
	if err != nil {
		return h.responser.Error(c, "useCase.GetWhere()", err)
	}

	return c.JSON(h.responser.OK(u))
}

func (h handler) GetAll(c echo.Context) error {
	users, err := h.useCase.GetAll()
	if err != nil {
		return h.responser.Error(c, "useCase.GetAll()", err)
	}

	return c.JSON(h.responser.OK(users))
}

func (h handler) GetByID(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.responser.Error(c, "uuid.Parse()", err)
	}

	user, err := h.useCase.GetByID(ID)
	if err != nil {
		return h.responser.Error(c, "useCase.GetWhere()", err)
	}

	return c.JSON(h.responser.OK(user))
}

func (h handler) Update(c echo.Context) error {
	var m model.User
	var err error

	m.ID, err = uuid.Parse(c.Param("id"))
	if err != nil {
		return h.responser.Error(c, "uuid.Parse()", err)
	}

	if err := c.Bind(&m); err != nil {
		return h.responser.BindFailed(err)
	}

	if file, err := c.FormFile("file"); err == nil {
		m.File = file
	}

	err = h.useCase.Update(&m)
	if err != nil {
		if strings.Contains(err.Error(), "the id does not exist") {
			resp := model.MessageResponse{
				Data:     "wrong ID",
				Messages: model.Responses{{Code: response.AuthError, Message: "wrong ID"}},
			}
			return c.JSON(http.StatusBadRequest, resp)
		}
		return h.responser.Error(c, "useCase.Update()", err)
	}

	return c.JSON(h.responser.Updated(m))
}

// func (h handler) AdminUpdate(c echo.Context) error {
// 	var m model.User
// 	var err error

// 	m.ID, err = uuid.Parse(c.Param("id"))
// 	if err != nil {
// 		return h.responser.Error(c, "uuid.Parse()", err)
// 	}

// 	if err := c.Bind(&m); err != nil {
// 		return h.responser.BindFailed(err)
// 	}

// 	err = h.useCase.AdminUpdate(&m)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "the id does not exist") {
// 			resp := model.MessageResponse{
// 				Data:     "wrong ID",
// 				Messages: model.Responses{{Code: response.AuthError, Message: "wrong ID"}},
// 			}
// 			return c.JSON(http.StatusBadRequest, resp)
// 		}
// 		return h.responser.Error(c, "useCase.Update()", err)
// 	}

//		return c.JSON(h.responser.Updated(m))
//	}
func (h handler) GetImage(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	img, err := h.useCase.GetImage(ID)
	if err != nil {
		return h.responser.Error(c, "useCase.GetImage()", err)
	}

	return c.File(img)
}
