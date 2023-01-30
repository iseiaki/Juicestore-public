package main

import (
	// "fmt"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
			"pr": "Add a product",
	})
}

func additempost(c *gin.Context) {
	name := c.PostForm("name")
		description := c.PostForm("description")
		nprice := c.PostForm("price")
		price, _  := strconv.Atoi(c.PostForm("price"))
		image := c.PostForm("image")
        if description != "" && nprice != "" && image != "" && name != "" {
		db.Create(&Product{Name: name, Description: description, Price: price, Image: image})
		}
		c.Redirect(http.StatusMovedPermanently, "/additem")
}

func deleteProduct(c *gin.Context){
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
    id := c.Param("id")
    realid, _ := strconv.Atoi(id)
    fmt.Println(realid)
    db.Where("id = ?", realid).Delete(&Product{})

}


func edititem (c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	var products []Product
		db.Find(&products)
		c.HTML(http.StatusOK, "edititems.html", gin.H{
			"products": products,
		})
}


func updateitem(c *gin.Context) {
	id := c.Param("id")
	productID, _ := strconv.Atoi(id)
	name := c.Param("name")
	description := c.Param("description")
	priceStr := c.Param("price")
	price, _ := strconv.Atoi(strings.TrimSuffix(priceStr, "$"))
	var product Product
	db.First(&product, productID)
	product.Name = name
	product.Description = description
	product.Price = price
	db.Save(&product)
}
