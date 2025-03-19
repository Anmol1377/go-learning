package main

import (
	"fmt"
	"strings"
)

func Variable() {
	// types of declaratin
	var name1 string = "1anmol"
	var name2 = "2name"
	name := "name3 "
	fmt.Println(name, name1, name2)

	// multi declatrion
	var n1, n2, n3, n4 = "1", "2", "3", "4"
	fmt.Println(n1, n2, n3, n4)
}

func stringextra() {
	//string erro
	// name := "my
	// dcd
	// d
	// d "
	// fmt.Println(name)

	// but we can use ``
	test_name := `gg
	g
	g
	g
	gg
	g
	`
	fmt.Println(test_name)
}

func string_code() {
	// string unicode with [] in go

	st := "anmol"

	fmt.Println(st[0])
	// %c is unicode decoder and printf is used for format prinitng
	fmt.Printf("%c", st[0])

	fmt.Println(len(st))

	str1, str2 := "ione", "ione"

	fmt.Println(strings.Compare(str1, str2))
}

func ifelse() {
	// condition
	urname := "an"

	if len(urname) < 5 {
		fmt.Println("its small dude")
	} else {
		fmt.Println("hmmm")
	}
}

func forloop() {
	// looping
	fmt.Println("Table of  2")
	for q := 1; q < 11; q++ {

		fmt.Println(q + q)
	}
}

func switch_code() {
	// switch
	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week!")
	case "Tuesday":
		fmt.Println("Second day of the week")
	case "Friday":
		fmt.Println("Almost weekend!")
	default:
		fmt.Println("Midweek day")
	}
}

func array_code() {
	// array
	arrq := [3]int{1, 2, 4}
	fmt.Println(arrq)

	// multidimension array
	var arr [2][3]int // A 2x3 array (2 rows, 3 columns)
	arr[0][0] = 1
	arr[0][1] = 2
	arr[0][2] = 3
	arr[1][0] = 4
	arr[1][1] = 5
	arr[1][2] = 6

	// Access values
	fmt.Println("array:", arr)

	// here we are initializating array
	var arr2 [3]string

	arr2[0] = "a"
	arr2[1] = "b"
	arr2[2] = "c"

	fmt.Println(arr2)

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	// extra ======
	arr3 := [3]int{10, 20, 30}

	for index, value := range arr3 {
		fmt.Println("Index:", index, "Value:", value)
	}

	// Ignore the index
	// The _ does not consume memory or resources—it simply tells Go to ignore the value.
	// The underscore (_) is a blank identifier in Go. It is used to ignore values you don't need or want to use.
	for _, value := range arr3 {
		fmt.Println("Value:", value)
	}

	for index := range arr3 {
		fmt.Println("Index:", index)
	}
}

func slice_code() {
	// slice

	slice := []int{}
	fmt.Println(slice)
	fmt.Println(cap(slice))

	var slik []int // Declare an empty slice

	slik = append(slik, 1)  // Capacity grows from 0 -> 1
	slik = append(slik, 3)  // Capacity grows from 1 -> 2
	slik = append(slik, 4)  // Capacity grows from 2 -> 4
	slik = append(slik, 6)  // Capacity grows from 4 -> 8
	slik = append(slik, 87) // No reallocation needed, slice is at capacity

	// it double at capacity end

	fmt.Println("Capacity:", cap(slik)) // Output: Capacity: 8

	// using make option for slice
	slik2 := make([]int, 5, 30)
	slik2[0] = 1
	slik2[1] = 3
	slik2[2] = 4
	slik2[3] = 6
	slik2[4] = 87

	slik2 = append(slik2, 3, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22, 2, 22, 2, 2)

	fmt.Println(slik2, "More data in slice with make keyword")

	// fmt.Println(len(slik2), "edew")

	// slice things
	myslc := []int{1, 2, 4, 5, 6}
	myslc = append(myslc, 1, 3, 34, 4, 4)
	fmt.Println(myslc[:])

	for i := 0; i < len(slik2); i++ {
		// took 0.8 sec for iteration of 414 element and some other task
		fmt.Println(slik2[i])
	}
}

func map_code() {
	// map data structure
	// name := map[type]type {}
	userinfo := map[string]string{

		"anmol": "nhi bataunga",
		"u":     "me",
		"test":  "test2",
	}

	userinfo["myname"] = "anmol"
	fmt.Println(userinfo)
	fmt.Println(userinfo["u"])

	for name, value := range userinfo {
		fmt.Println(name, ":", value)
	}

}

func struct_code() {
	// struct (structure )

	type worker struct {
		name   string
		age    int
		salary int
	}

	var people worker
	people.name = "anmol"
	people.age = 22
	people.salary = 1000

	fmt.Println(people)

	p2 := worker{
		name:   "anmol2",
		age:    11,
		salary: 2222,
	}

	fmt.Println(p2.name)
}

func number_print(number ...int) {
	fmt.Println(number)
	fmt.Println(cap(number))
}

func showname(s ...string) {
	fmt.Println(s)
}
