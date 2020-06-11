package handler

import (
	"AliveVirtualGift_RestAPI/src/proto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"AliveVirtualGift_RestAPI/src/connection"
	"AliveVirtualGift_RestAPI/src/middleware"
)

func Login(ctx *gin.Context) {

	var loginRequest proto.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	request := &loginRequest

	sc := connection.DialToAccountServiceServer()
	response, err := sc.ClientAccountService.Login(ctx, request)

	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusUnauthorized, "Invalid login details")
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "token",
		Value:    response.Token,
		HttpOnly: true,
	})
	ctx.JSON(http.StatusOK, response)
}

func Logout(ctx *gin.Context) {

	cookie, err := middleware.GetTokenCookie(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}
	tokenString := cookie.Value

	sc := connection.DialToAccountServiceServer()
	response, err := sc.ClientAccountService.Logout(ctx, &proto.LogoutRequest{
		Token: tokenString,
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "token",
		Value:    "",
		HttpOnly: true,
	})
	ctx.JSON(http.StatusOK, response)
}

func Create(ctx *gin.Context) {

}

func Update(ctx *gin.Context) {

}

func Show(ctx *gin.Context) {

}
