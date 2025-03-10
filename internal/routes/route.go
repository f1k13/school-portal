package routes

import (
	"github.com/gin-gonic/gin"
)

func StartRouter(r *gin.Engine) {

	AuthRouter(r)
}
