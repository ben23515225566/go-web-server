package main

import (
	"log"
	"net/http"
	"strconv"
	"web-server/database"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func signUp(c *gin.Context) {
	c.HTML(200, "signUp.html", gin.H{})
}

func signIn(c *gin.Context) {
	success := c.Query("success")
	c.HTML(200, "signIn.html", gin.H{
		"success": success,
	})
}

func articles(c *gin.Context) {
	success := c.Query("success")
	c.HTML(http.StatusOK, "list_articles.html", gin.H{
		"success": success,
	})
}

func article(c *gin.Context){
	c.HTML(http.StatusOK, "article.html", gin.H{})
}

func fetchArticles(c *gin.Context) {
	articles, err := database.QueryAllArticles()
	if err != nil {
		log.Panic("Query failed: ", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}

func fetchOneArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("Id"))
	article, err := database.QueryOneArticle(id)
	log.Println(article)
	if err != nil {
		log.Panic("Fetch one article: ", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"article": article,
	})
}

func InsertNewUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	err := database.InsertToUsers(username, password)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Insert new user: ", username)
	c.Redirect(http.StatusSeeOther, "/signIn?success=success") // 303: http.StatusSeeOther 用於重導向
}

func edit_article(c *gin.Context){
	c.HTML(http.StatusOK, "edit_article.html", gin.H{})
}

func create_new_article(c *gin.Context){
	title := c.PostForm("title")
	content := c.PostForm("content")
	err := database.InsertToArticles(title, content)
	if err != nil{
		log.Panic(err)
	}else{
		log.Println("Insert new article: ", title)
		c.Redirect(http.StatusSeeOther, "/articles?success=Create Successfully !")
	}

}

func delete_article(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("Id"))
	err := database.DeleteArticle(id)
	if err != nil{
		log.Panic(err)
	}else{
		log.Println("Delete article id: ", id)
		c.JSON(http.StatusOK, gin.H{
			"redirect": "103.179.29.154/articles",
		})
	}
}

func main() {
	defer database.CloseDatabase()

	router := gin.Default()
	router.LoadHTMLGlob("./template/html/*")
	router.Static("/assets", "./template/assets")

	api_router := router.Group("/api")
	{
		api_router.GET("/fetchArticles", fetchArticles)
		api_router.GET("/fetchOneArticle/:Id", fetchOneArticle)
		api_router.POST("/signUp", InsertNewUser)
		api_router.POST("/create_new_article", create_new_article)
		api_router.DELETE("/delete_article/:Id", delete_article)
	}

	view_router := router.Group("/")
	{
		view_router.GET("/", index)
		view_router.GET("/signUp", signUp)
		view_router.GET("/signIn", signIn)
		view_router.GET("/article/:id", article)
		view_router.GET("/articles", articles)
		view_router.GET("/edit_article", edit_article)
	}

	router.Run(":8080")
}
