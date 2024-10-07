package services

import (
	"time"

	"github.com/pramudyaws/synapsis-restaurant/config"
	"github.com/pramudyaws/synapsis-restaurant/models"
)

type PaymentTransactionService struct{}

func NewPaymentTransactionService() *PaymentTransactionService {
	return &PaymentTransactionService{}
}

func (s *PaymentTransactionService) SavePaymentTransaction(orderID uint, amountPaid float64) error {
	now := time.Now()
	paymentTransaction := models.PaymentTransaction{
		OrderID:    orderID,
		AmountPaid: amountPaid,
		PaidAt:     &now,
	}

	if err := config.DB.Create(&paymentTransaction).Error; err != nil {
		return err
	}
	return nil
}
