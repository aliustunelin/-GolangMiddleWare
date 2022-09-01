package main

import (
	"log"
	"net/http"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "book-worm-api/docs"

	"github.com/gin-gonic/gin"
)

// @version 1.0
// @description Tutorial
// @termsOfService http://swagger.io/terms/

// @contact.name Ali USTUNEL
// @contact.email aliustunelin@gmail.com
// @contact.url https://www.linkedin.com/in/ali-ustunel3338/

// @BasePath /api/v1

// @host localhost:5003
func main() {
	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")
	quotes := v1.Group("/quote")

	quotes.GET("/", GetAll)

	quotes.POST("/", Create)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	_ = router.Run(":5003")
}

// Create godoc
// @Summary Yeni bir kitap alıntısı ekler
// @Produce json
// @Accept json
// @Param quote body quote true "Alıntı Bilgileri"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /quote/ [post]
func Create(c *gin.Context) {
	var q quote

	if err := c.ShouldBindJSON(&q); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := AddQuote(&q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	q.ID = id
	c.JSON(http.StatusOK, gin.H{"added": q})
}

/*
   {array} is a quote list garranty
*/

// GetAll godoc
// @Summary Tüm kitap alıntılarını döndürür
// @Produce json
// @Success 200 {array} quote
// @Failure 500
// @Router /quote/ [get]
func GetAll(c *gin.Context) {

	var quoteList, err = GetQuotes()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"quotes": quoteList}) // HTTP 200 OK ile birlikte çekilen listeyi geri yolluyoruz.
}
