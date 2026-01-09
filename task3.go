package main

import (
	"fmt"
	"strings"
)

// 1. Функция encryptWord шифрует одно слово:
// Первый символ остается, остальные реверсируются.
func encryptWord(word string) string {
	// Преобразуем в []rune для корректной работы с Unicode (например, кириллицей)
	runes := []rune(word)

	if len(runes) <= 1 {
		return word
	}

	// Срез, который нужно перевернуть (все символы кроме первого)
	toReverse := runes[1:]

	// Классический алгоритм разворота массива/среза
	for i, j := 0, len(toReverse)-1; i < j; i, j = i+1, j-1 {
		toReverse[i], toReverse[j] = toReverse[j], toReverse[i]
	}

	// runes[0] остался на месте, а runes[1:] мы перевернули in-place
	return string(runes)
}

// 2. Функция encryptPhrase разбивает фразу, шифрует слова и собирает обратно.
func encryptPhrase(phrase string) string {
	words := strings.Fields(phrase)
	encryptedWords := make([]string, len(words))

	for i, w := range words {
		encryptedWords[i] = encryptWord(w)
	}

	return strings.Join(encryptedWords, " ")
}

func main() {
	// Тестовые фразы
	phrases := []string{
		"Pepe Schnele is a legend",
		"Hello World",
		"Go is awesome",
		"A simple test",
		"Кириллица тоже работает", // Проверка Unicode
	}

	fmt.Println("=== Шифратор фраз ===")
	for _, p := range phrases {
		encrypted := encryptPhrase(p)
		fmt.Printf("Исходная: \"%s\"\n", p)
		fmt.Printf("Шифр:     \"%s\"\n", encrypted)
		fmt.Println("---")
	}
}
