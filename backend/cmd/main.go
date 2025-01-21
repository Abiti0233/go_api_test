package main

import (
	"log"
	"net/http"
	"todo-app/interface/handler"

	"github.com/kelseyhightower/envconfig"

	"github.com/Abiti0233/go_api_test/backend/config"
	"github.com/Abiti0233/go_api_test/backend/infrasructure"
	"github.com/Abiti0233/go_api_test/backend/interface/handler"
	"github.com/Abiti0233/go_api_test/backend/usecase"
)

func main() {
	var cfg config.Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal("failed to process env var: %v", err)
	}

	// DB接続
	db, err := infrasructure.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("failed to initialize db: %v", err)
	}
	defer db.Close()

	// リポジトリ
	todoRepo := infrasructure.NewPostgresTodoRepository(db)

	// Usecase
	todoUC := usecase.NewTodoUseCase(todoRepo)

	// Handler
	todoHandler := handler.NewTodoHandler(todoUC)

	// ルーティング
	r := handler.NewRouter(todoHandler)

	// サーバー起動
	log.Printf("Starting server on port %s...", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
			log.Fatal(err)
	}
}