package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	auth "jd_workout_golang/app/middleware"
	authAction "jd_workout_golang/app/services/auth"
)

func RegisterUser(r *gin.RouterGroup) {
	fmt.Println("RegisterUser")
	testA()
	testB()
	r.POST("/register", authAction.RegisterAction)
	r.POST("/login", authAction.LoginAction)
	r.GET("/login/google", authAction.LoginWithGoogleAction)
	r.GET("/login/google/redirect", authAction.LoginWithGoogleAuthAction)
	r.POST("/login/google/access-token", authAction.LoginWithGoogleAccessTokenAction)

	userGroup := r.Group("/user").Use(auth.ValidateToken)

	userGroup.GET("/", authAction.InfoAction)
}

func testA() {
	fmt.Println("testA")
}

func testB() {
	fmt.Println("testB")
}

func testC() {
	fmt.Println("testC")
}

func testD() {
	fmt.Println("testD")
}
