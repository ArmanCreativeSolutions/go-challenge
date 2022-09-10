package routes

import (
	"github.com/gin-gonic/gin"
	user_segments "go-challenge/application/user-segments"
	"go-challenge/infrastructure/dbconfig"
	"go-challenge/infrastructure/repository"
	"go-challenge/presentation/usersegmentcontroller"
	"log"
)

const (
	prefix   = "/api/v1/"
	public   = ""
	cronJobs = "/cron-jobs"
)

func SetupRouter(r *gin.Engine) {

	dbService := dbconfig.NewSqliteDBService()
	_, err := dbService.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	userSegmentRepository := repository.NewUserSegmentRepository(dbService)
	userSegmentService := user_segments.NewUserSegmentService(userSegmentRepository)
	userSegmentController := usersegmentcontroller.NewUserSegmentController(userSegmentService)

	routes := r.Group(prefix)
	{
		publicApi := routes.Group(public)
		{
			publicApi.POST("user-segments", userSegmentController.CreateUserSegment)
			publicApi.GET("user-segments/count-segments", userSegmentController.CountSegmentsByTitle)
		}
		cronJobApi := routes.Group(cronJobs)
		{
			cronJobApi.POST("user-segments/remove-segment", userSegmentController.RemoveSegmentCronJob)
		}
	}
}
