package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	// Загружаем .env в процессе запуска (локально)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env не найден — пропускаем загрузку")
	}
}

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		fmt.Println("❌ TELEGRAM_TOKEN не найден")
		os.Exit(1)
	}
	fmt.Println("✅ TELEGRAM_TOKEN успешно получен:")
	fmt.Println(token)
}
