package main

import (
	"fmt"
	"go1f/pkg/api"
	"go1f/pkg/db"
	"go1f/pkg/server"
	"log"
)

func main() {

	// Создаем базу данных
	dbFile := "scheduler.db"

	if err := db.Init(dbFile); err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}
	log.Println("База данных успешно инициализирована.")

	// Иницилизируем Api
	api.Init()

	// Запускаем сервер
	log.Println("Запуск сервера.")
	if err := server.Run(); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %v", err)
	}

}
