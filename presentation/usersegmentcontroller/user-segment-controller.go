package usersegmentcontroller

import (
	"github.com/gin-gonic/gin"
	user_segments "go-challenge/application/user-segments"
	"go-challenge/application/user-segments/dto"
	"net/http"
)

type UserSegmentController struct {
	userSegmentService *user_segments.UserSegmentService
}

func NewUserSegmentController(userSegmentService *user_segments.UserSegmentService) *UserSegmentController {
	return &UserSegmentController{userSegmentService: userSegmentService}
}

func (*UserSegmentController) FetchUserSegments(c *gin.Context) {
	var createReq dto.CreateUserSegmentDTO
	err := c.ShouldBindJSON(createReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

func (*UserSegmentController) CreateUserSegment(c *gin.Context) {

}

func (*UserSegmentController) FetchUserSegmentById(c *gin.Context) {

}

func (*UserSegmentController) UpdateUserSegment(c *gin.Context) {

}

func (*UserSegmentController) DeleteUserSegment(c *gin.Context) {

}

func (*UserSegmentController) CountSegmentsByTitle(c *gin.Context) {

}

func (*UserSegmentController) RemoveSegmentScheduled(c *gin.Context) {

}
