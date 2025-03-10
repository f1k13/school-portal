package routes

import (
	userHandlers "github.com/f1k13/school-portal/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {
	r.POST("/auth/sign-up", userHandlers.SignUp)
}
