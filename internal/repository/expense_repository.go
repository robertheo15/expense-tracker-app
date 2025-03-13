package repository

import (
	"encoding/json"
	"errors"
	"expense-tracker-app/internal/model"
	"os"
	"sync"
)

type ExpenseRepository interface {
	LoadExpenses() ([]model.Expense, error)
	SaveExpenses(expenses []model.Expense) error
	Add(expense model.Expense) (int, error)
	GetAll() ([]model.Expense, error)
	Delete(id int) error
}

type ExpenseRepositoryImpl struct {
	filePath string
	mutex    sync.Mutex
}

func NewExpenseRepository(filePath string) ExpenseRepository {
	return &ExpenseRepositoryImpl{filePath: filePath}
}

// Load expenses from file
func (r *ExpenseRepositoryImpl) LoadExpenses() ([]model.Expense, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	file, err := os.Open(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Expense{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var expenses []model.Expense
	if err := json.NewDecoder(file).Decode(&expenses); err != nil {
		return nil, err
	}
	return expenses, nil
}

// Save expenses to file
func (r *ExpenseRepositoryImpl) SaveExpenses(expenses []model.Expense) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	file, err := os.Create(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(expenses)
}

// Add a new expense
func (r *ExpenseRepositoryImpl) Add(expense model.Expense) (int, error) {
	expenses, err := r.LoadExpenses()
	if err != nil {
		return 0, err
	}

	expense.ID = 1
	if len(expenses) > 0 {
		expense.ID = expenses[len(expenses)-1].ID + 1
	}

	expenses = append(expenses, expense)
	if err := r.SaveExpenses(expenses); err != nil {
		return 0, err
	}

	return expense.ID, nil
}

// Get all expenses
func (r *ExpenseRepositoryImpl) GetAll() ([]model.Expense, error) {
	return r.LoadExpenses()
}

// Delete an expense by ID
func (r *ExpenseRepositoryImpl) Delete(id int) error {
	expenses, err := r.LoadExpenses()
	if err != nil {
		return err
	}

	found := false
	newExpenses := []model.Expense{}
	for _, e := range expenses {
		if e.ID == id {
			found = true
		} else {
			newExpenses = append(newExpenses, e)
		}
	}

	if !found {
		return errors.New("expense not found")
	}

	return r.SaveExpenses(newExpenses)
}
