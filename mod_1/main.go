package main

import (
	"fmt"
	"reflect"
	"sort"
)

type HealthCheck struct {
	ServiceId int
	status    string
}

const (
	PasStatus  = "pass"
	FailStatus = "fail"
)

func GenerateCheck() (a []HealthCheck) {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			a = append(a, HealthCheck{i, PasStatus})
		} else {
			a = append(a, HealthCheck{i, FailStatus})
		}
	}
	return a
}

func NumsRepCheck(arr []int) bool {
	m := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				m++
				break
			}
		}
	}
	if m >= 2 {
		return true
	}
	return false
}

func SortSliceCheck(a []string) bool {
	b := make([]string, len(a))
	copy(b, a)
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	if reflect.DeepEqual(a, b) {
		return true
	}
	return false
}

func CountQuantityRep(s string) {
	runes := []rune(s)
	m := make(map[rune]int)
	for _, key := range runes {
		m[key]++
	}
	for key, value := range m {
		fmt.Println(string(key), " - ", value)
	}
}

func main() {
	//Task 1
	fmt.Println("\tTask 1")
	fmt.Println("Тут будет выведен идентификатор")
	a := GenerateCheck()
	for _, id := range a {
		if id.status == "pass" {
			fmt.Println(id.ServiceId)
		}
	}

	//Task2
	fmt.Println("\tTask 2")
	var arr = []int{2, 1, 1, 1, 4, 5, 6}
	fmt.Println(NumsRepCheck(arr))
	fmt.Println()

	//Task3
	fmt.Println("\tTask 3")
	l := []string{"C", "A", "B", "C"}
	//sort.Slice(l, func(i, j int) bool { return l[i] < l[j] })
	fmt.Println(SortSliceCheck(l))

	//Task4
	fmt.Println("\tTask 4")
	strCheck := "съешь еще этих мягких французких булок, да выпей чаю"
	fmt.Println(strCheck)
	CountQuantityRep(strCheck)
}
