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

	// バリデータの初期化
	adminValidator := validator.NewAdminValidator()
	wordBookValidator := validator.NewWordBookValidator()
	wordValidator := validator.NewWordValidator()
	difficultyValidator := validator.NewDifficultyValidator()

	// リポジトリの初期化
	adminRepository := repository.NewAdminRepository(dbConn)
	wordBookRepository := repository.NewWordBookRepository(dbConn)
	wordRepository := repository.NewWordRepository(dbConn)
	difficultyRepository := repository.NewDifficultyRepository(dbConn)
	genreRepository := repository.NewGenreRepository(dbConn)

	// ユースケースの初期化
	adminUsecase := usecase.NewAdminUsecase(adminRepository, adminValidator)
	wordBookUseCase := usecase.NewWordBookUsecase(wordBookRepository, wordBookValidator)
	wordUseCase := usecase.NewWordUsecase(wordRepository, wordValidator)
	difficultyUsecase := usecase.NewDifficultyUsecase(difficultyRepository, difficultyValidator)
	genreUsecase := usecase.NewGenreUsecase(genreRepository)

	// コントローラの初期化
	csrfTokenController := controller.NewCsrfTokenController()
	adminController := controller.NewAdminController(adminUsecase)
	wordBookController := controller.NewWordBookController(wordBookUseCase)
	wordController := controller.NewWordController(wordUseCase)
	difficultyController := controller.NewDifficultyController(difficultyUsecase)
	genreController := controller.NewGenreController(genreUsecase)

	// ルーターの初期化
	e := router.NewRouter(
		csrfTokenController,
		adminController,
		wordBookController,
		wordController,
		difficultyController,
		genreController,
	)

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}
