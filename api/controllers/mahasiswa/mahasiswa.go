package mahasiswa

import (
	"context"
	"golang-websocket/api/database"
	"golang-websocket/api/helper"
	"golang-websocket/api/models"
	"golang-websocket/api/repository/mahasiswa"
	"golang-websocket/api/usecase"
	ucase "golang-websocket/api/usecase/mahasiswa"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type MahasiswaHandler struct {
	MahasiswaUsecase usecase.MahasiswaUsecase
}

func NewMahasiswaHandler() MahasiswaHandler {
	timeout := time.Duration(viper.GetInt(`context.timeout`)) * time.Second
	db := database.Load()
	repoMahasiswa := mahasiswa.NewMahasiswaRepository(db)
	ucaseMahasiswa := ucase.NewMahasiswaUsecase(repoMahasiswa, timeout)
	return MahasiswaHandler{
		MahasiswaUsecase: ucaseMahasiswa,
	}
}

func (m *MahasiswaHandler) List(c *gin.Context) {
	var res = c.Writer
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	mahasiswa, err := m.MahasiswaUsecase.List(ctx)
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Responses(res, http.StatusOK, "Success", mahasiswa)
}

func (m *MahasiswaHandler) Detail(c *gin.Context) {
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

	mahasiswa, err := m.MahasiswaUsecase.Detail(ctx, id)
	if err != nil {
		helper.HandlerErrorQuery(res, err)
		return
	}
	helper.Responses(res, http.StatusOK, "Success", mahasiswa)
}

func (m *MahasiswaHandler) Insert(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	var res = c.Writer
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
		return
	}

	// var validators *validator.Validate
	// config := &validator.Config{TagName: "validate"}
	// validators = validator.New(config)
	// err := validators.Struct(mahasiswa)

	// if err != nil {
	// 	helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
	// 	return
	// }

	result, err := m.MahasiswaUsecase.Insert(ctx, mahasiswa)
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Responses(res, http.StatusOK, "Success", result)
}

func (m *MahasiswaHandler) Update(c *gin.Context) {
	var datas = make(map[string]interface{})
	var res = c.Writer
	id, err := helper.ToInt(c.Param("id"))
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
		return
	}

	datas["nama"] = c.Request.FormValue("nama")
	datas["nim"] = c.Request.FormValue("nim")
	datas["kelas"] = c.Request.FormValue("kelas")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	mahasiswa, err := m.MahasiswaUsecase.Update(ctx, datas, id)
	if err != nil {
		helper.HandlerErrorQuery(res, err)
		return
	}

	helper.Responses(res, http.StatusOK, "Succes", mahasiswa)
}

func (m *MahasiswaHandler) Delete(c *gin.Context) {
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

	err = m.MahasiswaUsecase.Delete(ctx, id)
	if err != nil {
		helper.HandlerErrorQuery(res, err)
		return
	}

	helper.Responses(res, http.StatusOK, "Success", "Data Telah Dihapus")
}
