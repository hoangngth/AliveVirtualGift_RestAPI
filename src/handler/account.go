package handler

import (
	"AliveVirtualGift_RestAPI/src/proto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"AliveVirtualGift_RestAPI/src/connection"
	"AliveVirtualGift_RestAPI/src/middleware"
)

//Login ...
func Login(ctx *gin.Context) {

	var loginRequest proto.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	sc := connection.DialToAccountServiceServer()
	response, err := sc.ClientAccountService.Login(ctx, &loginRequest)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "token",
		Value:    response.Token,
		HttpOnly: true,
	})

	ctx.JSON(http.StatusOK, response)
}

//Logout ...
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

//Create ...
func Create(ctx *gin.Context) {

	var createRequest proto.CreateRequest
	if err := ctx.ShouldBindJSON(&createRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	sc := connection.DialToAccountServiceServer()
	response, err := sc.ClientAccountService.Create(ctx, &createRequest)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

//Update ...
func Update(ctx *gin.Context) {

	var updateRequest proto.UpdateRequest
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	sc := connection.DialToAccountServiceServer()
	response, err := sc.ClientAccountService.Update(ctx, &updateRequest)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

//Show ...
func Show(ctx *gin.Context) {

	cookie, err := middleware.GetTokenCookie(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}
	tokenString := cookie.Value

	sc := connection.DialToAccountServiceServer()
	response, err := sc.ClientAccountService.Show(ctx, &proto.ShowRequest{
		Token: tokenString,
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
