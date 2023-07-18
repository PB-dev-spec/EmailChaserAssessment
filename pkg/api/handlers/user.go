package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"mark_emailchaser/pkg/api/responses"
	"mark_emailchaser/pkg/models"
	"mark_emailchaser/pkg/service"
	"mark_emailchaser/pkg/utils"
)

func CreateUser(ctx *gin.Context) {
	// Parse the request and validate the input
	var newUser models.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		responses.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Perform more validations
	if len(newUser.Password) < 8 {
		responses.ErrorResponse(ctx, http.StatusBadRequest, "password should be at least 8 characters long")
		return
	}

	if !utils.IsEmailValid(newUser.Email) {
		responses.ErrorResponse(ctx, http.StatusBadRequest, "invalid email format")
		return
	}

	// Hash the user password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		responses.ErrorResponse(ctx, http.StatusInternalServerError, "error while hashing the password")
		return
	}
	newUser.Password = string(hashedPassword)

	// Call the user service to create a new user in the database
	user, err := service.CreateUser(newUser)
	if err != nil {
		responses.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	// Clear password before returning user
	user.Password = ""
	ctx.JSON(http.StatusCreated, gin.H{"user": user})
}

func Login(ctx *gin.Context) {
	var credentials models.User
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		responses.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := service.GetUserByEmail(credentials.Email)
	if err != nil {
		responses.ErrorResponse(ctx, http.StatusUnauthorized, "invalid email or password")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		fmt.Println(err) // Print the error
		responses.ErrorResponse(ctx, http.StatusUnauthorized, "invalid email or password")
		return
	}
	// At this point, the user is authenticated. Now we create a JWT.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Hour * 72).Unix(), // Token expires after 72 hours
	})
	// Replace "secret" with your secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		responses.ErrorResponse(ctx, http.StatusInternalServerError, "could not login")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func GetProfile(ctx *gin.Context) {
	userIDStr := ctx.Param("id")
	fmt.Println(userIDStr)
	// Convert the user ID string to an integer.
	userIDInt, err := strconv.Atoi(userIDStr)
	if err != nil {
		// If the conversion failed, return an error response.
		responses.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID format.")
		return
	}
	userID := uint(userIDInt)
	// Use the UserService to fetch the user from the database.
	user, err := service.GetUserByID(userID)
	if err != nil {
		// If there was an error fetching the user, return an error response.
		responses.ErrorResponse(ctx, http.StatusInternalServerError, "Error fetching user profile.")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func DeleteUser(ctx *gin.Context) {
	userIDStr := ctx.Param("id")
	// Convert the user ID string to an integer.
	userIDInt, err := strconv.Atoi(userIDStr)
	if err != nil {
		// If the conversion failed, return an error response.
		responses.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID format.")
		return
	}
	userID := uint(userIDInt)
	// Use the UserService to delete the user from the database.
	err = service.DeleteUser(userID)
	if err != nil {
		// If there was an error deleting the user, return an error response.
		responses.ErrorResponse(ctx, http.StatusInternalServerError, "Error deleting user.")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
