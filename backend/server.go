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
	wordBookUseCase := usecase.NewWordBookUsecase(wordBookRepository)
	wordBookController := controller.NewWordBookController(wordBookUseCase)
	e := router.NewRouter(wordBookController)

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}
