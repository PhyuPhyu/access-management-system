package controllers

import (
	"access-management-system/models"
	"access-management-system/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
)

// Serve health check
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Server is up and running",
	})
}

// User register
func UserRegister(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(user.Password) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Password should not be empty!"})
		return
	}

	err := models.GenerateAndSetPasswordHash(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.CreateUser(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate Verification Code
	code := randstr.String(20)

	verificationCode := utils.Encode(code)

	// Update code in db
	user.EmailVerificationCode = verificationCode
	err = models.UpdateUser(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send email verification code to user
	err = utils.SendEmailVerificationCode(user, code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Email sent successfully to " + user.Email})
}

// Email Verification by user
func VerifyEmail(c *gin.Context) {
	code := c.Params.ByName("verification_code")
	verificationCode := utils.Encode(code)

	var user models.User
	err := models.GetUserByEmailVerifyCode(&user, verificationCode)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.EmailVerificationCode = ""
	user.EmailVerified = true
	err = models.UpdateUser(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}

// User login
func UserLogin(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err := models.GetUserByEmail(&user, input.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Input email is not registered yet"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	if !user.EmailVerified {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Please verify your email"})
		return
	}

	token, err := utils.CreateToken(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User login successfully", "token": token})
}

// Get user profile
func GetMe(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(models.User)

	response := &models.UserResponse{
		ID:    currentUser.ID,
		Name:  currentUser.Name,
		Email: currentUser.Email,
	}

	c.JSON(http.StatusOK, response)
}

// ..........................................................................Admin....................................................................
// Get all user lists
func GetAllUsers(c *gin.Context) {
	var users []models.User
	err := models.GetUsers(&users)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Get user by id
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	err := models.GetUser(&user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// User account approved by admin
func UserAccountApprovedByAdmin(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	var user models.User
	err := models.GetUser(&user, userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.IsAllowByAdmin {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Already approved"})
		return
	}

	if user.EmailVerified {
		user.IsAllowByAdmin = true

		err = models.UpdateUser(&user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Send admin approval email to user
		err = utils.SendApprovedEmail(user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "This account requires email verification"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User account approved successfully"})
}

// Delete user by admin
func DeleteUser(c *gin.Context) {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
