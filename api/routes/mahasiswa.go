package routes

import (
	"database/sql"
	handler "golang-websocket/api/controllers/mahasiswa"
	repo "golang-websocket/api/repository/mahasiswa"
	ucase "golang-websocket/api/usecase/mahasiswa"
	"time"

	"github.com/gin-gonic/gin"
)

func RouteMahasiswa(route *gin.RouterGroup, db *sql.DB, timeOut time.Duration) {
	repoMahasiswa := repo.NewMahasiswaRepository(db)
	ucaseMahasiswa := ucase.NewMahasiswaUsecase(repoMahasiswa, timeOut)
	handlerMahasiswa := handler.MahasiswaHandler{MahasiswaUsecase: ucaseMahasiswa}

	router := route.Group("/mahasiswa")
	{
		router.GET("/list", handlerMahasiswa.List)
		router.GET("/detail/:id", handlerMahasiswa.Detail)
		router.POST("/insert", handlerMahasiswa.Insert)
		router.PUT("/update/:id", handlerMahasiswa.Update)
		router.DELETE("/delete/:id", handlerMahasiswa.Delete)
	}
}
