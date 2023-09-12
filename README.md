
# ginErrorHandler

`ginErrorHandler` is a Go package that provides a global error handler middleware for the Gin web framework. It collects error messages from Gin errors and responds with a JSON payload containing the error messages.

## Installation

To install this package, use the standard `go get` command:

``` shell

go get -u github.com/abdelrhman-basyoni/ginErrorHandler

```

## Usage

### Middleware Setup

To use `ginErrorHandler` as a middleware in your Gin application, follow these steps:

1. Import the package:

   ```go
   
   import "github.com/yourusername/ginErrorHandler"
   

   ```

2. Create a Gin router instance:

   ```go
   
   r := gin.Default()
   
   ```

3. Add the `GlobalErrorHandler` middleware to your Gin router  :

   ```go
   // add it first thing under the r:= gin.Default()
   r := gin.Default()
   r.Use(ginErrorHandler.GlobalErrorHandler("errors"))
   
   ```

4. Define your Gin routes and handlers as usual.

5. Run your Gin application:

   ```go
   r.Run(":8080")
   ```
6. Inside your controller when you handle error handle it like this
```go

	token, err := useCases.Login(login.Email, login.Password)
	if err != nil {
		ginContext.Error(err)
		return
	}
```

### Customization

You can customize the error key used in the JSON response by passing it as an argument to the `GlobalErrorHandler` function.

```go
r.Use(ginErrorHandler.GlobalErrorHandler("customErrorKey"))
```



## Contributing

Contributions to `ginErrorHandler` are welcome. If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.



## Contact

For questions or feedback, please contact:

- Email: abdelrhman.mbasyoni@gmail.com
- GitHub: [GitHub Profile](https://github.com/abdelrhman-basyoni)

## Acknowledgments

Special thanks to the Gin web framework community and contributors.

