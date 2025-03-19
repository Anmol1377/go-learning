package main

// create url shortner for go project

// import "fmt"
// import "time"

import (
	"fmt"
	"time"
)

type PaymentGateway interface {
	ProcessPayment(amount float64) string
	RefundPayment(transactionID string) string
}

type payment struct {
	gateway PaymentGateway
}

type PayPal struct{}

func (p PayPal) ProcessPayment(amount float64) string {
	return fmt.Sprintf("PayPal: Processed payment of $%.2f", amount)
}

func (p PayPal) RefundPayment(transactionID string) string {
	return fmt.Sprintf("PayPal: Refunded payment for transaction %s", transactionID)
}

type Stripe struct{}

func (s Stripe) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Stripe: Processed payment of $%.2f", amount)
}

func (s Stripe) RefundPayment(transactionID string) string {
	return fmt.Sprintf("Stripe: Refunded payment for transaction %s", transactionID)
}

func main() {
	start := time.Now()

	// variale
	// variable()

	// // string extra
	// stringextra()

	// // string operatopn with unicode
	// string_code()

	// // if else
	// ifelse()

	// // for loop
	// forloop()

	// // swithc code
	// switch_code()

	// // array code
	// array_code()

	// // slice
	// slice_code()

	// // map code
	// map_code()

	// // struct code
	// struct_code()

	// number print
	// number_print(1, 3, 4, 4, 9, 44)

	// show user
	// showname("ghgh", "hjghg")

	// retuning funtion
	// add := func(numa int, numb int) int {
	// 	return numa + numb
	// }
	// fmt.Println(add(34, 4))

	// imidiate invoked funtion
	// func() {
	// 	fmt.Println("heloo for inf")
	// }()

	// Create a Person instance
	person := Person{FirstName: "John", LastName: "Doe"}

	// Call the method on the Person instance
	fullName := person.FullName()

	fmt.Println(person.FirstName)
	fmt.Println("Full Name:", fullName)

	// callback
	myCallback := func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}

	greet("Anmol", myCallback)

	// defer code here will move hoi1 to last of func
	mydelaycode()

	// pointet
	add := 20
	// creting pointer var
	var pointer *int = &add

	// ,memory address here
	fmt.Println(&add)
	fmt.Println(*pointer)

	// panic funtion below
	empName := "saam"
	age := 75
	employee(&empName, age)

	// generic

	t1, t2 := showitm("12hg", "20")
	fmt.Println(t1, t2)

	// todo interface=====
	// interface
	// Create an instance of Person
	var speaker Speaker = Person2{Name: "John"}

	// Call the Speak method from the interface
	fmt.Println(speaker.Speak())

	// chanels

	ch := make(chan int)

	go sendData(ch)

	value := <-ch

	fmt.Println("Received:", value)

	// interface here

	paymentdone := payment{
		gateway: PayPal{},
	}
	fmt.Println(paymentdone.gateway.ProcessPayment(12121222))

	elapsed := time.Since(start)                                                      // Calculate elapsed time
	fmt.Printf("Program executed in %s\n", elapsed)                                   // Original output
	fmt.Printf("Program executed in %.3f ms\n", float64(elapsed.Microseconds())/1000) // In milliseconds

}

// method named "FullName" for the type "Person"
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

type Person struct {
	FirstName string
	LastName  string
}

// callback funtion
func greet(name string, callback func(string)) {
	callback(name)
}

// defer keyword
func mydelaycode() {
	defer fmt.Println("hoi1")
	fmt.Println("hoi2")
	fmt.Println("hoi3")
	fmt.Println("hoi4")
	fmt.Println("hoi5")
}

// panic func
func employee(name *string, age int) {
	if age > 65 {
		// panic("Age cannot be greater than retirement age")
	}

}

// generic
func showitm[T any](an, bn T) (T, T) {
	return an, bn
}

// interface code
// Define the Speaker interface
type Speaker interface {
	Speak() string
}

// Define a struct for Person
type Person2 struct {
	Name string
}

// Implement the Speak method for Person
func (p Person2) Speak() string {
	return p.Name + " says hello!"
}

func sendData(ch chan<- int) {
	ch <- 42
}
