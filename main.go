package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

func main() {
	// Парсинг флагов
	pathFlag := flag.String("path", "", "Путь к файлу с URL-адресами (относительный или абсолютный)")
	sizeFlag := flag.String("size", "bytes", "Единицы измерения размера контента: bytes, kb, mb, chars")
	saveFlag := flag.Bool("save", false, "Сохранить результат в файл (true/false)")

	flag.Parse()

	// Проверка на обязательные параметры
	if *pathFlag == "" {
		log.Fatal("Флаг --path обязателен")
	}

	// Определение единицы измерения
	var sizeUnit SizeUnit
	switch strings.ToLower(*sizeFlag) {
	case "kb":
		sizeUnit = KB
	case "mb":
		sizeUnit = MB
	case "chars":
		sizeUnit = Chars
	case "bytes":
		sizeUnit = Bytes
	default:
		log.Fatalf("Неверное значение для --size: %s. Используйте bytes, kb, mb или chars", *sizeFlag)
	}

	// Обработка файла
	links, err := ProcessFile(*pathFlag)
	if err != nil {
		log.Fatalf("Ошибка обработки файла: %v", err)
	}
	if len(links) == 0 {
		log.Fatal("В файле не найдено валидных URL-адресов")
	}
	fmt.Println("Обработка урлов закончена.")

	// Обработка сайтов
	results := ProcessSites(links, sizeUnit)
	fmt.Println("Обработка сайтов закончена.")

	// Вывод результатов
	DisplayResults(results, sizeUnit)

	// Сохранение результатов, если требуется
	if *saveFlag {
		outputFile := filepath.Join(filepath.Dir(*pathFlag), "results.txt")
		err := SaveResultsToFile(results, sizeUnit, outputFile)
		if err != nil {
			log.Fatalf("Ошибка сохранения результатов в файл: %v", err)
		}
		fmt.Printf("Результаты сохранены в файл: %s\n", outputFile)
	}
}
