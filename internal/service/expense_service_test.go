package service_test

import (
	"expense-tracker-app/internal/model"
	"expense-tracker-app/internal/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock Repository
type MockExpenseRepository struct {
	mock.Mock
}

func (m *MockExpenseRepository) LoadExpenses() ([]model.Expense, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Expense), args.Error(1)
}

func (m *MockExpenseRepository) SaveExpenses(expenses []model.Expense) error {
	args := m.Called(expenses)
	return args.Error(0)
}

func (m *MockExpenseRepository) Add(expense model.Expense) (int, error) {
	args := m.Called(expense)
	return args.Int(0), args.Error(1)
}

func (m *MockExpenseRepository) GetAll() ([]model.Expense, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Expense), args.Error(1)
}

func (m *MockExpenseRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestAddExpense(t *testing.T) {
	mockRepo := new(MockExpenseRepository)
	svc := service.NewExpenseService(mockRepo)

	expense := model.Expense{
		Description: "Lunch",
		Amount:      15.50,
		Date:        time.Now(),
	}

	mockRepo.On("Add", mock.Anything).Return(1, nil)

	id, err := svc.AddExpense(expense.Description, expense.Amount)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)
	mockRepo.AssertExpectations(t)
}

func TestGetAllExpenses(t *testing.T) {
	mockRepo := new(MockExpenseRepository)
	svc := service.NewExpenseService(mockRepo)

	expenses := []model.Expense{
		{ID: 1, Description: "Lunch", Amount: 10.50, Date: time.Now()},
		{ID: 2, Description: "Dinner", Amount: 20.00, Date: time.Now()},
	}

	mockRepo.On("GetAll").Return(expenses, nil)

	result, err := svc.GetAllExpenses()
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestDeleteExpense(t *testing.T) {
	mockRepo := new(MockExpenseRepository)
	svc := service.NewExpenseService(mockRepo)

	mockRepo.On("Delete", 1).Return(nil)

	err := svc.DeleteExpense(1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSummary(t *testing.T) {
	mockRepo := new(MockExpenseRepository)
	svc := service.NewExpenseService(mockRepo)

	expenses := []model.Expense{
		{ID: 1, Description: "Lunch", Amount: 10.50, Date: time.Date(2025, 3, 10, 0, 0, 0, 0, time.UTC)},
		{ID: 2, Description: "Dinner", Amount: 20.00, Date: time.Date(2025, 3, 15, 0, 0, 0, 0, time.UTC)},
	}

	mockRepo.On("GetAll").Return(expenses, nil)

	total, err := svc.Summary(3) // Fix: use `svc.Summary(3)` instead of `service.Summary(3)`
	assert.NoError(t, err)
	assert.Equal(t, 30.50, total)
	mockRepo.AssertExpectations(t)
}
