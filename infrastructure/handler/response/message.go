package response

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/mariobac1/backend_webpages/model"
)

const (
	BindFailed      = "bind_failed"
	Ok              = "ok"
	RecordCreated   = "record_created"
	RecordUpdated   = "record_updated"
	RecordDeleted   = "record_deleted"
	UnexpectedError = "unexpected_error"
	AuthError       = "authorization_error"
	FieldsEmpty     = "fields_empty_error"
)

type API struct{}

// New returns a new response API
func New() API {
	return API{}
}

func (a API) OK(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: Ok, Message: "¡listo!"}},
	}
}

func (a API) Created(data interface{}) (int, model.MessageResponse) {
	return http.StatusCreated, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: RecordCreated, Message: "¡listo!"}},
	}
}

func (a API) Updated(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: RecordUpdated, Message: "¡listo!"}},
	}
}

func (a API) Image() (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     "Ok",
		Messages: model.Responses{{Code: RecordUpdated, Message: "¡listo!"}},
	}
}

func (a API) Deleted(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: RecordDeleted, Message: "¡listo!"}},
	}
}

func (a API) BindFailed(err error) error {
	message := "BindFailed"
	if strings.Contains(err.Error(), "the header is empty") || strings.Contains(err.Error(), "the token is not valid") || strings.Contains(err.Error(), "you are not admin") {
		message = "You don't have authorization"
	}
	e := model.NewError()
	e.Err = err
	e.APIMessage = message
	e.Code = BindFailed
	e.StatusHTTP = http.StatusBadRequest
	e.Who = "c.Bind()"

	log.Warnf("%s", e.Error())
	return &e
}

func (a API) Error(c echo.Context, who string, err error) *model.Error {
	var message string
	var code string
	var status int

	switch {
	case strings.Contains(err.Error(), "Fields can't be empty"):
		{
			message = "The fields can't be empty"
			code = FieldsEmpty
			status = http.StatusBadRequest
		}
	default:
		{
			message = "¡Sorry! we have internal problems"
			code = UnexpectedError
			status = http.StatusInternalServerError
		}
	}
	e := model.NewError()
	e.Err = err
	e.APIMessage = message
	e.Code = code
	e.StatusHTTP = status
	e.Who = who

	userID, ok := c.Get("userID").(uuid.UUID)
	// Only to avoid the panic error
	if !ok {
		log.Errorf("cannot get/parse uuid from userID")
	}
	e.UserID = userID.String()

	log.Errorf("%s", e.Error())
	return &e
}
