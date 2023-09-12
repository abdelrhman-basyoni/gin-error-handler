package ginErrorHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomParserType func(errors []string) []string

func GlobalErrorHandler(customErrorParser CustomParserType) gin.HandlerFunc {
	var messages []string

	return func(c *gin.Context) {
		// Iterate through Gin errors and convert them to your custom Error type
		c.Next()

		for _, ginErr := range c.Errors {
			messages = append(messages, ginErr.Error())
		}

		errorsList := customErrorParser(messages)

		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"errors": &errorsList})
	}
}
