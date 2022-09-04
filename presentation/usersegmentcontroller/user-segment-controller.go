package usersegmentcontroller

import (
	"github.com/gin-gonic/gin"
	user_segments "go-challenge/application/user-segments"
	"go-challenge/application/user-segments/dto"
	Response "go-challenge/infrastructure/response"
	"go-challenge/infrastructure/validation"
	"log"
	"net/http"
	"strings"
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

	userSegmentResponse, responseError := userSegmentController.userSegmentService.CreateUserSegment(createReq)

	if responseError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": Response.ErrorResponse{Error: responseError.Error()}})
		return
	}
	response := Response.GeneralResponse{
		Error:   false,
		Message: "user segment created",
		Data:    userSegmentResponse,
	}
	c.JSON(http.StatusOK, gin.H{"response": response})
}

func (*UserSegmentController) FetchUserSegmentById(c *gin.Context) {

}

func (*UserSegmentController) UpdateUserSegment(c *gin.Context) {

}

func (*UserSegmentController) DeleteUserSegment(c *gin.Context) {

}

func (userSegmentController *UserSegmentController) CountSegmentsByTitle(c *gin.Context) {
	segmentTitle := c.Query("title")
	segmentTitle = strings.TrimSpace(segmentTitle)
	if segmentTitle == "" {
		response := Response.GeneralResponse{
			Error:   true,
			Message: "segment title is empty",
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	countSegmentsResponse, err := userSegmentController.userSegmentService.CountUserSegmentsBySegmentTitle(segmentTitle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": Response.ErrorResponse{Error: err.Error()}})
		return
	}
	response := Response.GeneralResponse{
		Error:   false,
		Message: "",
		Data:    countSegmentsResponse,
	}
	c.JSON(http.StatusOK, gin.H{"response": response})
}

func (*UserSegmentController) RemoveSegmentScheduled(c *gin.Context) {

}
