package routes

import (
	"github.com/gin-gonic/gin"
	user_segments "go-challenge/application/user-segments"
	"go-challenge/domain/usersegments"
	"go-challenge/presentation/usersegmentcontroller"
)

const (
	prefix = "/api/v1/"
	public = ""
)

func SetupRouter(r *gin.Engine) {

	userSegmentRepository := usersegments.NewUserSegmentRepository()
	userSegmentService := user_segments.NewUserSegmentService(userSegmentRepository)
	userSegmentController := usersegmentcontroller.NewUserSegmentController(userSegmentService)

	routes := r.Group(prefix)
	{
		publicApi := routes.Group(public)
		{
			publicApi.GET("user-segments", userSegmentController.FetchUserSegments)
			publicApi.POST("user-segments", userSegmentController.CreateUserSegment)
			publicApi.GET("user-segments/:id", userSegmentController.FetchUserSegmentById)
			publicApi.PUT("user-segments/:id", userSegmentController.UpdateUserSegment)
			publicApi.DELETE("user-segments/:id", userSegmentController.DeleteUserSegment)
			publicApi.GET("user-segments/count-segments/:title", userSegmentController.CountSegmentsByTitle)
			publicApi.POST("user-segments/remove-segment-scheduled", userSegmentController.RemoveSegmentScheduled)
		}
	}
}
