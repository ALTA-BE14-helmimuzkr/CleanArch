package main

import (
	"api/config"
	bookData "api/features/book/data"
	bookHandler "api/features/book/handler"
	bookService "api/features/book/services"
	userData "api/features/user/data"
	userHandler "api/features/user/handler"
	userService "api/features/user/services"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	userData := userData.New(db)
	userSrv := userService.New(userData)
	userHdl := userHandler.New(userSrv)

	bookData := bookData.New(db)
	bookSrv := bookService.New(bookData)
	bookHdl := bookHandler.New(bookSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom}, method=${method}, uri=${uri}, status=${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())

	e.GET("/users/profile", userHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/users", userHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/users", userHdl.Deactive(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/books", bookHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/books/:id", bookHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/books/:id", bookHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/books/mybook", bookHdl.MyBook(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/books", bookHdl.GetAllBook())

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
