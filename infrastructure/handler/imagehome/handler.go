package imagehome

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/mariobac1/backend_webpages/domain/imagehome"
	"github.com/mariobac1/backend_webpages/infrastructure/handler/response"
	"github.com/mariobac1/backend_webpages/model"
)

type handler struct {
	useCase   imagehome.UseCase
	responser response.API
}

func newHandler(uc imagehome.UseCase) handler {
	return handler{useCase: uc}

}

func (h handler) Create(c echo.Context) error {
	var m model.ImageHome

	if err := c.Bind(&m); err != nil {
		if strings.Contains(err.Error(), "the header is empty") {
			resp := model.MessageResponse{
				Data:     "the header is empty",
				Messages: model.Responses{{Code: response.AuthError, Message: "You don't have authorization"}},
			}
			return c.JSON(http.StatusBadRequest, resp)
		}
		return h.responser.BindFailed(err)
	}

	if file, err := c.FormFile("file"); err == nil {
		m.File = file
	}

	// file, err := c.FormFile("file")
	// if err != nil {
	// 	return h.responser.Error(c, "FormFile()", err)
	// }
	// m.File = file
	if err := h.useCase.Create(&m); err != nil {
		return h.responser.Error(c, "useCase.Create()", err)
	}

	return c.JSON(h.responser.Created(m))
}

func (h handler) GetAll(c echo.Context) error {
	imagehomes, err := h.useCase.GetAll()
	if err != nil {
		return h.responser.Error(c, "useCase.GetAll()", err)
	}

	return c.JSON(h.responser.OK(imagehomes))
}

func (h handler) GetByID(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.responser.Error(c, "uuid.Parse()", err)
	}

	imageHomeData, err := h.useCase.GetByID(ID)
	if err != nil {
		return h.responser.Error(c, "useCase.GetWhere()", err)
	}

	return c.JSON(h.responser.OK(imageHomeData))
}

func (h handler) Update(c echo.Context) error {
	var m model.ImageHome
	var err error
	fmt.Printf("Este es el contexto que viene %v", c)
	fmt.Printf("Este es el valor que viene %v", m)
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

func (h handler) GetImage(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	img, err := h.useCase.GetImage(ID)
	if err != nil {
		return h.responser.Error(c, "useCase.GetImage()", err)
	}

	return c.File(img)
}
