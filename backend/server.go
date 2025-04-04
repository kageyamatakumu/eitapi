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
	difficultyRepository := repository.NewDifficultyRepository(dbConn)

	adminUsecase := usecase.NewAdminUsecase(adminRepository, adminValidator)
	wordBookUseCase := usecase.NewWordBookUsecase(wordBookRepository, wordBookValidator)
	wordUseCase := usecase.NewWordUsecase(wordRepository, wordValidator)
	difficultyUsecase := usecase.NewDifficultyUsecase(difficultyRepository)

	csrfTokenController := controller.NewCsrfTokenController()
	adminController := controller.NewAdminController(adminUsecase)
	wordBookController := controller.NewWordBookController(wordBookUseCase)
	wordController := controller.NewWordController(wordUseCase)
	difficultyController := controller.NewDifficultyController(difficultyUsecase)

	e := router.NewRouter(
		csrfTokenController,
		adminController,
		wordBookController,
		wordController,
		difficultyController,
	)

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}
