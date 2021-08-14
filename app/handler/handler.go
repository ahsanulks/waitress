package handler

import "github.com/gin-gonic/gin"

// render for set response api.
func render(c *gin.Context, body interface{}, status int) {
	switch v := body.(type) {
	case string:
		c.JSON(status, buildMessageResponse(body.(string), status))
	case error:
		c.JSON(status, struct {
			Error string `json:"error"`
		}{
			Error: v.Error(),
		})
	case nil:
		c.Status(status)
	default:
		c.JSON(status, buildSuccessResponse(body, status))
	}
}

type MessageResponse struct {
	Message string     `json:"message"`
	Meta    HTTPStatus `json:"http_status"`
}

func buildMessageResponse(message string, status int) MessageResponse {
	response := MessageResponse{
		Message: message,
		Meta: HTTPStatus{
			HTTPStatus: status,
		},
	}
	return response
}

func buildSuccessResponse(data interface{}, status int) SuccessResponse {
	return SuccessResponse{
		Data: data,
		Meta: HTTPStatus{
			HTTPStatus: status,
		},
	}
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
	Meta HTTPStatus  `json:"meta"`
}

type HTTPStatus struct {
	HTTPStatus int `json:"http_status"`
}
