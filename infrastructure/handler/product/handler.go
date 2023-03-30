package product

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/mariobac1/backend_webpages/domain/product"
	"github.com/mariobac1/backend_webpages/infrastructure/handler/response"
	"github.com/mariobac1/backend_webpages/model"
)

type handler struct {
	useCase   product.UseCase
	responser response.API
}

func newHandler(uc product.UseCase) handler {
	return handler{useCase: uc}

}

func (h handler) Create(c echo.Context) error {
	var m model.Product

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

	file, err := c.FormFile("file")
	if err != nil {
		return h.responser.Error(c, "FormFile()", err)
	}

	m.File = file

	if err := h.useCase.Create(&m); err != nil {
		return h.responser.Error(c, "useCase.Create()", err)
	}

	return c.JSON(h.responser.Created(m))
}

func (h handler) GetAll(c echo.Context) error {
	products, err := h.useCase.GetAll()
	if err != nil {
		return h.responser.Error(c, "useCase.GetAll()", err)
	}

	return c.JSON(h.responser.OK(products))
}

func (h handler) GetByID(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.responser.Error(c, "uuid.Parse()", err)
	}

	bookingData, err := h.useCase.GetByID(ID)
	if err != nil {
		return h.responser.Error(c, "useCase.GetWhere()", err)
	}

	return c.JSON(h.responser.OK(bookingData))
}

func (h handler) Update(c echo.Context) error {
	var m model.Product

	if err := c.Bind(&m); err != nil {
		return h.responser.BindFailed(err)
	}

	err := h.useCase.Update(&m)
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
