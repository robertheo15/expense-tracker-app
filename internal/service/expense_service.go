package service

import (
	"errors"
	"expense-tracker-app/internal/model"
	"expense-tracker-app/internal/repository"
	"time"
)

type ExpenseService interface {
	AddExpense(description string, amount float64) (int, error)
	GetAllExpenses() ([]model.Expense, error)
	DeleteExpense(id int) error
	Summary(month int) (float64, error)
}

type ExpenseServiceImpl struct {
	repo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) ExpenseService {
	return &ExpenseServiceImpl{repo: repo}
}

// Add an expense
func (s *ExpenseServiceImpl) AddExpense(description string, amount float64) (int, error) {
	if amount <= 0 {
		return 0, errors.New("amount must be positive")
	}

	expense := model.Expense{
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
	}

	return s.repo.Add(expense)
}

// Get all expenses
func (s *ExpenseServiceImpl) GetAllExpenses() ([]model.Expense, error) {
	return s.repo.GetAll()
}

// Delete an expense
func (s *ExpenseServiceImpl) DeleteExpense(id int) error {
	return s.repo.Delete(id)
}

func (s *ExpenseServiceImpl) Summary(month int) (float64, error) {
	expenses, err := s.repo.GetAll()
	if err != nil {
		return 0, err
	}

	var total float64
	currentYear := time.Now().Year()

	for _, e := range expenses {
		if e.Date.Year() == currentYear && (month == 0 || int(e.Date.Month()) == month) {
			total += e.Amount
		}
	}

	return total, nil
}
