package services

import (
	"api-kasir/models"
	"api-kasir/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []models.CheckoutItems) (*models.Transaction, error) {
	// Implementation for processing checkout will go here
	return s.repo.CreateTransaction(items)
}