package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"trigonal/backend-auth/api/database"
	"trigonal/backend-auth/api/helper"
	"trigonal/backend-auth/api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var body models.User

	fmt.Println("parse body data to variabel")
	if err := ctx.BindJSON(&body); err != nil {
		fmt.Println(err.Error())
		helper.ReturnValidateError(err, ctx)
		return
	}

	fmt.Println("hash user password")
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		helper.ERROR(ctx, http.StatusInternalServerError, "Gagal menyimpan data")
		return
	}
	fmt.Println("input hashed password to variable")
	body.Password = string(hash)

	fmt.Println("create data user")
	newUser := database.DB.Create(&body)
	if newUser.Error != nil {
		helper.ERROR(ctx, http.StatusInternalServerError, "Gagal menyimpan data")
		return
	}

	var user models.User
	database.DB.Omit("password").First(&user, body.ID)
	fmt.Println("return success create user")
	helper.SUCCESS(ctx, http.StatusCreated, user)
}

func Login(ctx *gin.Context)  {
	var body models.UserLogin

	fmt.Println("parse body data to variabel")
	if err := ctx.BindJSON(&body); err != nil {
		helper.ReturnValidateError(err, ctx)
		return
	}

	fmt.Println("check user in DB")
	var user models.User
	database.DB.Where("email = ?", body.EmailOrUsername).Or("username = ?", body.EmailOrUsername).First(&user)

	if user.ID == uuid.Nil {
		helper.ERROR(ctx, http.StatusBadRequest, "Akun tidak terdaftar")
		return
	}

	fmt.Println("comparing password in db and request body")
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password));err != nil {
		helper.ERROR(ctx, http.StatusBadRequest, "Kata sandi salah")
		return
	}

	fmt.Println("generate token for user")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": map[string]string{
			"ID": user.ID.String(),
		},
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY_API")))
	if err != nil {
		helper.ERROR(ctx, http.StatusInternalServerError, "Gagal masuk, coba lagi")
		return
	}

	fmt.Println("return success login")
	var userOutput models.User
	database.DB.Omit("password").First(&userOutput, user.ID)
	helper.SUCCESS(ctx, http.StatusOK, gin.H{
		"user": userOutput,
		"token": tokenString,
	})
}