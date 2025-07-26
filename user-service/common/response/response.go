package response

import (
	"net/http"

	"github.com/anddriii/warunggo/user-service/constants"
	errConst "github.com/anddriii/warunggo/user-service/constants/error"
	"github.com/gin-gonic/gin"
)

// format API response
type Response struct {
	Status  string      `json:"status"`
	Message any         `json:"any"`
	Data    interface{} `json:"data"`
	Token   *string     `json:"token,omitempty"` // jika token tidak diisi, maka "Token" tidak masuk ke json
}

type ParamHTTPResp struct {
	Code    int
	Err     error
	Message *string
	Gin     *gin.Context
	Data    interface{}
	Token   *string
}

func HttpResponse(param ParamHTTPResp) {
	if param.Err == nil {
		param.Gin.JSON(param.Code, Response{
			Status:  constants.Succes,
			Message: http.StatusText(http.StatusOK),
			Data:    param.Data,
			Token:   param.Token,
		})
		return
	}

	message := errConst.ErrInternalServerError.Error()
	if param.Message != nil {
		message = *param.Message
	} else if param.Err != nil {
		if errConst.ErrMapping(param.Err) {
			message = param.Err.Error()
		}
	}

	param.Gin.JSON(param.Code, Response{
		Status:  constants.Error,
		Message: message,
		Data:    param.Data,
	})
	return
}
