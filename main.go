package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")
    router.GET("/", mainpage)
    router.POST("/submit", additempost)
    router.DELETE("/delete/:id",deleteProduct)
    router.GET("/edititems", edititem)
    router.POST("/edititem/:id/:name/:description/:price", updateitem)
    router.GET("/additem", additem)
    router.Run()
}
