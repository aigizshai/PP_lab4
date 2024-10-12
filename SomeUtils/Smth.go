package someutils

import (
	"fmt"
	"strings"
)

func Upper() string {
	fmt.Println("Введите строку")
	var s string
	fmt.Scan(&s)
	str := strings.ToUpper(s)
	fmt.Println("Cтрока в верхнем регистре")
	fmt.Println(str)
	return str
}

func Sum() int {
	fmt.Println("Введите 3 числа")
	var a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Println("Сумма ", a+b+c)
	return a + b + c
}

func ReverseArr() [4]int {
	fmt.Print("Заполните массив")
	var a [4]int
	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println("Исходный массив")
	fmt.Println(a)

	for i := 0; i < len(a)/2; i++ {
		sw := a[len(a)-1-i]
		a[len(a)-1-i] = a[i]
		a[i] = sw
	}
	fmt.Println("Перевернутый массив")
	fmt.Println(a)

	return a
}
