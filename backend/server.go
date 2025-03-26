package main

import (
	"backend/controller"
	"backend/db"
	"backend/repository"
	"backend/router"
	"backend/usecase"
)

func main() {

	dbConn := db.CreateDB()
	adminRepository := repository.NewAdminRepository(dbConn)
	wordBookRepository := repository.NewWordBookRepository(dbConn)
	wordRepository := repository.NewWordRepository(dbConn)
	adminUsecase := usecase.NewAdminUsecase(adminRepository)
	wordBookUseCase := usecase.NewWordBookUsecase(wordBookRepository)
	wordUseCase := usecase.NewWordUsecase(wordRepository)
	adminController := controller.NewAdminController(adminUsecase)
	wordBookController := controller.NewWordBookController(wordBookUseCase)
	wordController := controller.NewWordController(wordUseCase)
	e := router.NewRouter(adminController, wordBookController, wordController)

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}
