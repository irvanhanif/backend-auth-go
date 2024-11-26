package controllers

import (
	"fmt"
	"net/http"
	"trigonal/backend-auth/api/database"
	"trigonal/backend-auth/api/helper"
	"trigonal/backend-auth/api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CheckUserAccess(ctx *gin.Context, user *models.User) (success bool) {
	var userToken models.User
	
	userInToken,_ := ctx.Get("user")
	userToken = userInToken.(models.User)
	
	if userToken.ID != user.ID {
		helper.ERROR(ctx, http.StatusUnauthorized, "Unauthorized user")
		return false
	}

	return true
}

func GetAllUsers(ctx *gin.Context) {
	var users []models.User

	database.DB.Omit("password").Find(&users)
	helper.SUCCESS(ctx, http.StatusOK, users)
}

func GetUser(ctx *gin.Context)  {
	var user models.User
	id := ctx.Param("id")

	database.DB.Omit("password").First(&user, map[string]string{"id": id})
	if user.ID == uuid.Nil {
		helper.ERROR(ctx, http.StatusNotFound, "User not found")
		return
	}

	if isOk := CheckUserAccess(ctx, &user); !isOk{
		return
	}

	helper.SUCCESS(ctx, http.StatusOK, user)
}

func UpdateUser(ctx *gin.Context)  {
	var body models.User
	var user models.User

	fmt.Println("parse request body to variable")
	if err := ctx.BindJSON(&body); err != nil {
		helper.ERROR(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("get data from db and check user exist")
	database.DB.First(&user, map[string]string{"id": body.ID.String()})
	if user.ID == uuid.Nil {
		helper.ERROR(ctx, http.StatusNotFound, "User not found")
		return
	}
	body.Password = user.Password
	
	if isOk := CheckUserAccess(ctx, &user); !isOk{
		return
	}

	fmt.Println("update data user")
	if err := database.DB.Model(&user).Updates(body).Error; err != nil {
		helper.ERROR(ctx, http.StatusInternalServerError, err.Error())
	}

	fmt.Println("return success update")
	helper.SUCCESS(ctx, http.StatusOK, "Updated user Successfully")
}

func DeleteUser(ctx *gin.Context)  {
	var user models.User
	id := ctx.Param("id")

	database.DB.First(&user, map[string]string{"id": id})
	if user.ID == uuid.Nil {
		helper.ERROR(ctx, http.StatusNotFound, "User not found")
		return
	}
	
	if isOk := CheckUserAccess(ctx, &user); !isOk{
		return
	}

	database.DB.Unscoped().Delete(&user)

	helper.SUCCESS(ctx, http.StatusOK, "Deleted user successfully")	
}