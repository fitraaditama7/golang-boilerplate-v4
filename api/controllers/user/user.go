package user

import (
	"context"
	"golang-websocket/api/database"
	"golang-websocket/api/helper"
	"golang-websocket/api/models"
	"golang-websocket/api/repository/user"
	"golang-websocket/api/usecase"
	ucase "golang-websocket/api/usecase/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler() UserHandler {
	timeout := time.Duration(viper.GetInt(`context.timeout`)) * time.Second
	db := database.Load()
	repoUser := user.NewUserRepository(db)
	ucaseUser := ucase.NewUserUsecase(repoUser, timeout)
	return UserHandler{
		UserUsecase: ucaseUser,
	}
}

func (u *UserHandler) List(c *gin.Context) {
	var res = c.Writer
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	users, err := u.UserUsecase.List(ctx)
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Responses(res, http.StatusOK, "Success", users)
}

func (u *UserHandler) Detail(c *gin.Context) {
	var res = c.Writer
	id, err := helper.ToInt(c.Param("id"))
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
		return
	}
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	user, err := u.UserUsecase.Detail(ctx, id)
	if err != nil {
		helper.HandlerErrorQuery(res, err)
		return
	}

	helper.Responses(res, http.StatusOK, "Success", user)
}

func (u *UserHandler) Insert(c *gin.Context) {
	var res = c.Writer
	var user models.User
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
		return
	}

	email := helper.IsEmail(user.Email)
	if email == false {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, "Invalid email")
		return
	}

	result, err := u.UserUsecase.Insert(ctx, user)
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Responses(res, http.StatusOK, "Success", result)
}

func (u *UserHandler) Update(c *gin.Context) {
	var res = c.Writer
	var datas = make(map[string]interface{})
	id, err := helper.ToInt(c.Param("id"))
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadGateway, err.Error())
		return
	}

	datas["nama"] = c.Request.FormValue("nama")
	datas["username"] = c.Request.FormValue("username")
	datas["password"] = c.Request.FormValue("password")
	datas["email"] = c.Request.FormValue("email")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := u.UserUsecase.Update(ctx, datas, id)
	if err != nil {
		helper.HandlerErrorQuery(res, err)
		return
	}
	helper.Responses(res, http.StatusOK, "Success", user)
}

func (u *UserHandler) Delete(c *gin.Context) {
	var res = c.Writer
	id, err := helper.ToInt(c.Param("id"))
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = u.UserUsecase.Delete(ctx, id)
	if err != nil {
		helper.HandlerErrorQuery(res, err)
		return
	}

	helper.Responses(res, http.StatusOK, "Success", "Data Telah Dihapus")
}
