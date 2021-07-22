package main

import (
	"embed"
	"html/template"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/j1mmyson/reviewList/controller"
	"github.com/j1mmyson/reviewList/models"
)

var (
	//go:embed web/templates/*
	templatesFS embed.FS
	//go:embed web
	staticFS embed.FS
)

func main() {

	r := gin.Default()
	LoadHTMLFromEmbedFS(r, templatesFS, "web/templates/*")

	r.GET("/static/*filepath", func(c *gin.Context) {
		c.FileFromFS(path.Join("/web/", c.Request.URL.Path), http.FS(staticFS))
	})

	models.ConnectDB()

	r.GET("/", controller.LogInPage)
	r.POST("/", controller.LogIn)
	r.GET("/signup", controller.SignUpPage)
	r.POST("/signup", controller.SignUp)
	r.GET("/logout", controller.LogOut)
	r.GET("/dashboard", controller.DashBoardPage)

	r.GET("/lists", controller.AllLists)
	r.POST("/lists", controller.CreateList)
	r.GET("/lists/:user", controller.FindListByUserName)
	r.POST("/delete/:id", controller.DeleteListById)
	r.POST("/edit/:id", controller.EditListById)

	r.Run(":8080")
}

func LoadHTMLFromEmbedFS(r *gin.Engine, em embed.FS, pattern string) {
	templ := template.Must(template.ParseFS(em, pattern))

	r.SetHTMLTemplate(templ)
}
