package ginErrorHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorType uint64
type Error struct {
	Err  error
	Type ErrorType
	Meta any
}

type ErrorMsgs []*Error

type CustomParserType func(errors ErrorMsgs) string

func GlobalErrorHandler(customErrorParser CustomParserType) gin.HandlerFunc {
	var messages ErrorMsgs

	return func(c *gin.Context) {
		// Iterate through Gin errors and convert them to your custom Error type
		for _, ginErr := range c.Errors {
			customErr := &Error{
				Err:  ginErr.Err,
				Type: ErrorType(ginErr.Type),
				Meta: ginErr.Meta,
			}
			messages = append(messages, customErr)
		}

		c.Next()
		errorsList := customErrorParser(messages)

		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"errors": &errorsList})
	}
}
