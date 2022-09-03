package routes

import (
	"github.com/gin-gonic/gin"
	"go-challenge/presentation/usersegmentcontroller"
)

func SetupRouter(r *gin.Engine) {
	r.GET("user-segments", usersegmentcontroller.FetchUserSegments)
	r.POST("users", usersegmentcontroller.CreateUserSegment)
	r.GET("users/:id", usersegmentcontroller.FetchUserSegmentById)
	r.PUT("users/:id", usersegmentcontroller.UpdateUserSegment)
	r.DELETE("users/:id", usersegmentcontroller.DeleteUserSegment)
}
