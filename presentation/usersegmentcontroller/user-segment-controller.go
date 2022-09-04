package usersegmentcontroller

import (
	"github.com/gin-gonic/gin"
	user_segments "go-challenge/application/user-segments"
	"go-challenge/application/user-segments/dto"
	Response "go-challenge/infrastructure/response"
	"go-challenge/infrastructure/validation"
	"log"
	"net/http"
)

type UserSegmentController struct {
	userSegmentService *user_segments.UserSegmentService
}

func NewUserSegmentController(userSegmentService *user_segments.UserSegmentService) *UserSegmentController {
	return &UserSegmentController{userSegmentService: userSegmentService}
}

func (*UserSegmentController) FetchUserSegments(c *gin.Context) {

}

func (userSegmentController *UserSegmentController) CreateUserSegment(c *gin.Context) {
	var createReq dto.CreateUserSegmentDTO
	err := c.ShouldBindJSON(&createReq)
	if err != nil {
		log.Println(err)
	}

	validationError := validation.ValidationCheck(createReq)
	if validationError != nil {
		response := Response.GeneralResponse{
			Error:   true,
			Message: validationError.Error(),
		}
		c.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	userResponse, responseError := userSegmentController.userSegmentService.CreateUserSegment(createReq)

	if responseError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": Response.ErrorResponse{Error: responseError.Error()}})
		return
	}
	response := Response.GeneralResponse{
		Error:   false,
		Message: "user segment created",
		Data:    userResponse,
	}
	c.JSON(http.StatusOK, gin.H{"response": response})
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
