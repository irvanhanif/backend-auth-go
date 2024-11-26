package helper

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func MsgForTag(tag string) string {
	switch tag {
	case "required":
		return "Kolom ini wajib diisi"
	case "email":
		return "Email belum sesuai"
	case "min":
		return "Jawaban kolom ini kurang panjang"
	}
	return ""
}

func ReturnValidateError(err error, ctx *gin.Context) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{fe.Field(), MsgForTag(fe.Tag())}
		}
		fmt.Println(err.Error())
		ERROR(ctx, http.StatusBadRequest, gin.H{
			"validate": out,
		})
	}
}