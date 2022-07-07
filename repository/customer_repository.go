package repository

import (
	"context"
	"fmt"
	"golang_lopei_grc_client/service"
	"log"
)

type CustomerRepository interface {
	CheckBalance(lopeiId int32) (float32, error)
	DoPayment(lopeiId int32, amount float32) error
}

type customerRepository struct {
	client service.LopeiPaymentClient
}

func (c *customerRepository) DoPayment(lopeiId int32, amount float32) error {
	payment, err := c.client.DoPayment(context.Background(), &service.PaymentMessage{
		LopeId: lopeiId,
		Amount: amount,
	})

	if err != nil {
		log.Fatalln(" Error when calling do payment...", err)
	}

	fmt.Println(payment)
	return nil 
}

func (c *customerRepository) CheckBalance(lopeiId int32)(float32, error) {
	balance, err := c.client.CheckBalance(context.Background(), &service.CheckBalanceMessage{LopeiId: lopeiId})

	if err != nil {
		log.Fatalln("Error when calling check balance...", err)
	}
	fmt.Println(balance)
	return 0, err
}

func NewCustomerRepository(client service.LopeiPaymentClient) CustomerRepository {
	repo := new(customerRepository)
	repo.client = client 
	return repo 
}