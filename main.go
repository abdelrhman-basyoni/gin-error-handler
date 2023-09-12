package ginErrorHandler

import (
	"fmt"
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
		c.Next()
		fmt.Print("custom handler works")
		for _, ginErr := range c.Errors {
			customErr := &Error{
				Err:  ginErr.Err,
				Type: ErrorType(ginErr.Type),
				Meta: ginErr.Meta,
			}
			messages = append(messages, customErr)
		}

		errorsList := customErrorParser(messages)

		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"errors": &errorsList})
	}
}
