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
	wordBookRepository := repository.NewWordBookRepository(dbConn)
	wordRepository := repository.NewWordRepository(dbConn)
	wordBookUseCase := usecase.NewWordBookUsecase(wordBookRepository)
	wordUseCase := usecase.NewWordUsecase(wordRepository)
	wordBookController := controller.NewWordBookController(wordBookUseCase)
	wordController := controller.NewWordController(wordUseCase)
	e := router.NewRouter(wordBookController, wordController)

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}
