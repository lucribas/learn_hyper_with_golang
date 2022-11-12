package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	// json tag to de-serialize json body
	Name string `json:"name"`
}

func main() {
	engine := gin.New()
	engine.POST("/test", func(context *gin.Context) {
		body := Body{}
		// using BindJson method to serialize body with struct
		if err := context.BindJSON(&body); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(body)
		body.Name = "Hello " + body.Name
		context.JSON(http.StatusAccepted, &body)
	})
	engine.Run(":3000")
}
