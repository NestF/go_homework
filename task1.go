package main

import (
	"fmt"
)

// 1. Создайте структуру PepeSchnele с тремя полями: Speed, Charisma, Wisdom
type PepeSchnele struct {
	Speed    int
	Charisma int
	Wisdom   int
}

// 2. Напишите функцию NewPepeSchnele... возвращает указатель на нового Пепе Шнеле
func NewPepeSchnele(speed, charisma, wisdom int) *PepeSchnele {
	return &PepeSchnele{
		Speed:    speed,
		Charisma: charisma,
		Wisdom:   wisdom,
	}
}

// 3. Напишите метод GetRating() для структуры PepeSchnele
func (p *PepeSchnele) GetRating() int {
	return (p.Speed * 2) + (p.Charisma * 3) + p.Wisdom
}

// 4. Реализация String() для форматированного вывода
// "Пепе Шнеле [Скорость: X, Харизма: Y, Мудрость: Z] | Рейтинг: R"
func (p *PepeSchnele) String() string {
	return fmt.Sprintf("Пепе Шнеле [Скорость: %d, Харизма: %d, Мудрость: %d] | Рейтинг: %d",
		p.Speed, p.Charisma, p.Wisdom, p.GetRating())
}

func main() {

	// Создайте как минимум двух разных Пепе Шнеле
	pepe1 := NewPepeSchnele(10, 5, 8) // Rating: 10*2 + 5*3 + 8 = 20 + 15 + 8 = 43
	pepe2 := NewPepeSchnele(7, 9, 10) // Rating: 7*2 + 9*3 + 10 = 14 + 27 + 10 = 51

	// Выведите информацию о каждом с помощью fmt.Println
	fmt.Println(pepe1)
	fmt.Println(pepe2)

	// Определите, у кого рейтинг выше
	rating1 := pepe1.GetRating()
	rating2 := pepe2.GetRating()

	if rating1 > rating2 {
		fmt.Println("Первый Пепе Шнеле круче!")
	} else if rating2 > rating1 {
		fmt.Println("Второй Пепе Шнеле круче!")
	} else {
		fmt.Println("Оба Пепе Шнеле равны по силе!")
	}
}
