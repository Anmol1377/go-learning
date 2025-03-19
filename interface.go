package main

import "fmt"

// PaymentGateway interface defines methods for processing payments
type PaymentGateway interface {
    ProcessPayment(amount float64) string
    RefundPayment(transactionID string) string
}


type payment struct {
    gateway PaymentGateway
}

// PayPal struct implements the PaymentGateway interface
type PayPal struct{}

func (p PayPal) ProcessPayment(amount float64) string {
    return fmt.Sprintf("PayPal: Processed payment of $%.2f", amount)
}

func (p PayPal) RefundPayment(transactionID string) string {
    return fmt.Sprintf("PayPal: Refunded payment for transaction %s", transactionID)
}

// Stripe struct implements the PaymentGateway interface
type Stripe struct{}

func (s Stripe) ProcessPayment(amount float64) string {
    return fmt.Sprintf("Stripe: Processed payment of $%.2f", amount)
}

func (s Stripe) RefundPayment(transactionID string) string {
    return fmt.Sprintf("Stripe: Refunded payment for transaction %s", transactionID)
}

func main() {
    var gateway PaymentGateway

    paymentdone := payment{
        gateway: PayPal{}
    } 

    fmt.Println(paymentdone.ProcessPayment(12121222));


    // Using PayPal
    // gateway = PayPal{}
    // fmt.Println(gateway.ProcessPayment(100.0))
    // fmt.Println(gateway.RefundPayment("TX123"))

    // // Using Stripe
    // gateway = Stripe{}
    // fmt.Println(gateway.ProcessPayment(200.0))
    // fmt.Println(gateway.RefundPayment("TX456"))
}
