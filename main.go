package main

import "fmt"

type Employee struct { //Employee - працівник
	Name     string
	Age      int
	Salary   int //зарплата
	FullTime bool
}

func main() {
	Jack := Employee{
		Name:     "Jack Smith",
		Age:      27,
		Salary:   4000,
		FullTime: false,
	}
	Jill := Employee{
		Name:     "Jill Jones",
		Age:      41,
		Salary:   60000,
		FullTime: true,
	}

	var employees []Employee
	employees = append(employees, Jack)
	employees = append(employees, Jill)

	for _, x := range employees {
		if x.Age > 30 {
			fmt.Println(x.Name, "is 30 or older")
		} else {
			fmt.Println(x.Name, "is under 30")
		}

		if x.Age > 30 && x.Salary > 50000 {
			fmt.Println(x.Name, "makes more than 50000 and is over 30")
		} else {
			fmt.Println("Either", x.Name, "makes less than 50000 or is under 30")
		}

		if x.Age > 30 || x.Salary > 50000 {
			fmt.Println(x.Name, "makes more than 50000 or is over 30")
		} else {
			fmt.Println(x.Name, "makes less than 50000 and is under 30")
		}
		if x.Age > 30 || x.Salary < 50000 && x.FullTime {
			fmt.Println(x.Name, "mathes our unclear criteria")
		}
		if (x.Age > 30 || x.Salary < 50000) && x.FullTime {
			fmt.Println(x.Name, "mathes our unclear criteria")
		}
	}
}
