package responses

import "github.com/labstack/echo/v4"

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Data struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Response(context echo.Context, statusCode int, data interface{}) error {
	return context.JSON(statusCode, data)
}

func ErrorResponse(context echo.Context, statusCode int, message string) error {
	return Response(context, statusCode, Error{
		Code:  statusCode,
		Error: message,
	})
}

func MessageResponse(context echo.Context, statusCode int, message string) error {
	return Response(context, statusCode, Data{
		Code:    statusCode,
		Message: message,
	})
}
