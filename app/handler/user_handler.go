package handler

import (
	"hr-management-system/app/request"
	"hr-management-system/app/response"
	"hr-management-system/app/service"
	"hr-management-system/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
	authService service.AuthService
}

func NewUserHandler(userService service.UserService, authService service.AuthService) *userHandler {
	return &userHandler{userService, authService}
}

// @Tags Users
// @Summary Login User
// @Description Login User
// @Accept  json
// @Produce  json
// @Param input body request.LoginInput true "Login"
// @Success 200
// @Router /login [post]
func (h *userHandler) Login(c *gin.Context) {
	var input request.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.APIResponse("Login failed", http.StatusUnprocessableEntity, false, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := utils.APIResponse("Login failed", http.StatusUnprocessableEntity, false, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := utils.APIResponse("Login failed", http.StatusBadRequest, false, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := response.ResponseLogin(loggedinUser, token)
	response := utils.APIResponse("Successfuly Loggedin", http.StatusOK, true, formatter)

	c.JSON(http.StatusOK, response)
}
