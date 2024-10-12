package workwithmaps

import "fmt"

func CreateMap() map[string]int {
	m := make(map[string]int)
	m["Алекскей"] = 5
	m["Слава"] = 20

	return m
}

func AddNew(m map[string]int) map[string]int {
	fmt.Println("Введите имя")
	var s string
	fmt.Scan(&s)
	fmt.Println("Введите возраст")
	var age int
	fmt.Scan(&age)
	m[s] = age
	return m
}

func AverAge(m map[string]int) float32 {
	var age float32
	for i, _ := range m {
		age += float32(m[i])
	}

	return age / float32(len(m))
}

func Del(m map[string]int) map[string]int {
	fmt.Println(m)
	fmt.Println("Введите имя для удаления")
	var s string
	fmt.Scan(&s)
	delete(m, s)

	fmt.Println("После измемнения")
	fmt.Println(m)

	return m
}
