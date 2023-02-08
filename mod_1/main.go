package main

import "fmt"

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

func main() {
	fmt.Println("Тут будет выведен идентификатор")
	a := GenerateCheck()
	for _, id := range a {
		if id.status == "pass" {
			fmt.Println(id.ServiceId)
		}
	}
}
