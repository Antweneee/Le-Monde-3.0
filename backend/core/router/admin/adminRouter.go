package articles

import (
	"github.com/gin-gonic/gin"
)

func ApplyAdminRoutes(public *gin.RouterGroup) {

	public.POST("/register", Register)

	public.POST("/login", Login)
}
