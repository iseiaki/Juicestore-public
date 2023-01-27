package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func mainpage(c *gin.Context){
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
		var products []Product
		db.Find(&products)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"products": products,
	})
}

func additem(c *gin.Context){
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

		c.HTML(http.StatusOK, "additem.html", gin.H{
			"pr": "Add an item",
	})
}

func additempost(c *gin.Context) {
	name := c.PostForm("name")
		description := c.PostForm("description")
		price, _  := strconv.Atoi(c.PostForm("price"))
		image := c.PostForm("image")

		db.Create(&Product{Name: name, Description: description, Price: price, Image: image})
		c.Redirect(http.StatusMovedPermanently, "/additem")
}

func deleteProduct(c *gin.Context){
    id := c.Param("id")
	realid, _  := strconv.Atoi(id)
	fmt.Println(realid)
	db.Where("id = ?", realid).Delete(&Product{})
    c.Redirect(http.StatusMovedPermanently, "/")
}