package docs

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterSwagger(v1 gin.IRouter) {
	url := ginSwagger.URL("http://localhost:8080/api/v1/swagger/doc.json")
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
