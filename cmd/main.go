package main


import (
	"github.com/gin-gonic/gin"

	"github.com/mhmdryhn/cognito-auth-service/route"
)


func main()  {
	router := gin.Default()
	route.RegisterRoutes(router)
	router.Run(":8080")
}
