package handler

import (
	"log"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
)

func InitializeStripe(apiKey string) {
	stripe.Key = apiKey
}
func CreateCustomer(email string) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{
		Email: stripe.String(email),
	}

	customer, err := customer.New(params)
	if err != nil {
		log.Printf("Failed to create customer: %v\n", err)
		return nil, err
	}

	return customer, nil
}
func ChargeCustomer(customerID string, amount int64, currency string) (*stripe.Charge, error) {
	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
		Customer: stripe.String(customerID),
	}

	charge, err := charge.New(params)
	if err != nil {
		log.Printf("Failed to charge customer: %v\n", err)
		return nil, err
	}

	return charge, nil
}
