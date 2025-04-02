package main

import (
	"backend/controller"
	"backend/db"
	"backend/repository"
	"backend/router"
	"backend/usecase"
	"backend/validator"
)

func main() {

	dbConn := db.CreateDB()

	adminValidator := validator.NewAdminValidator()
	wordBookValidator := validator.NewWordBookValidator()
	wordValidator := validator.NewWordValidator()

	adminRepository := repository.NewAdminRepository(dbConn)
	wordBookRepository := repository.NewWordBookRepository(dbConn)
	wordRepository := repository.NewWordRepository(dbConn)

	adminUsecase := usecase.NewAdminUsecase(adminRepository, adminValidator)
	wordBookUseCase := usecase.NewWordBookUsecase(wordBookRepository, wordBookValidator)
	wordUseCase := usecase.NewWordUsecase(wordRepository, wordValidator)

	adminController := controller.NewAdminController(adminUsecase)
	wordBookController := controller.NewWordBookController(wordBookUseCase)
	wordController := controller.NewWordController(wordUseCase)

	e := router.NewRouter(adminController, wordBookController, wordController)

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}
