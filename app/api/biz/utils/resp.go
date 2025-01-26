package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	c.JSON(code, &ErrorResponse{
		StatusCode: 500,
		StatusMsg:  err.Error(),
	})
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}
