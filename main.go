// Package ginErrorHandler provides a global error handler middleware for the Gin web framework.
// It collects error messages from Gin errors and responds with a JSON payload containing
// the error messages.
package ginHandler

// GlobalErrorHandler is a middleware for Gin that collects error messages from Gin errors
// and responds with a JSON payload containing the error messages.
//
// Parameters:
//
//	errorKey (string): The key to use in the JSON response to represent the error messages.
//
// Example:
//
//	r := gin.Default()
//	r.Use(ginErrorHandler.GlobalErrorHandler("errors"))
//	// ...
//	r.Run(":8080")

func GlobalErrorHandler(errorKey string) HandlerFunc {
	var messages []string
	var statusCode = 500
	return func(c *Context) {
		// Iterate through Gin errors and convert them to your custom Error type
		c.Next()

		for _, ginErr := range c.Errors {
			messages = append(messages, ginErr.Error())
			if ginErr.code != nil {
				statusCode = *ginErr.code
			}

		}

		c.AbortWithStatusJSON(statusCode, map[string]interface{}{errorKey: &messages})
	}
}
