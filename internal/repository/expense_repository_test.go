package repository_test

import (
	"encoding/json"
	"expense-tracker-app/internal/model"
	"expense-tracker-app/internal/repository"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setupTempFile(t *testing.T) string {
	tmpfile, err := os.CreateTemp("", "expenses_*.json")
	if err != nil {
		t.Fatal(err)
	}

	// Write an empty JSON array to prevent EOF issues
	_, err = tmpfile.Write([]byte("[]"))
	if err != nil {
		t.Fatal(err)
	}

	tmpfile.Close()
	return tmpfile.Name()
}

func setupRepoWithData(t *testing.T, tempFile string, expenses []model.Expense) {
	data, err := json.Marshal(expenses)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(tempFile, data, 0644); err != nil {
		t.Fatal(err)
	}
}

func TestAddExpense(t *testing.T) {
	tempFile := setupTempFile(t)
	defer os.Remove(tempFile)

	repo := repository.NewExpenseRepository(tempFile)

	expense := model.Expense{
		Description: "Coffee",
		Amount:      5.99,
		Date:        time.Now(),
	}

	id, err := repo.Add(expense)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)
}

func TestGetAllExpenses(t *testing.T) {
	tempFile := setupTempFile(t)
	defer os.Remove(tempFile)

	expenses := []model.Expense{
		{ID: 1, Description: "Coffee", Amount: 5.99, Date: time.Now()},
		{ID: 2, Description: "Lunch", Amount: 12.50, Date: time.Now()},
	}
	setupRepoWithData(t, tempFile, expenses)

	repo := repository.NewExpenseRepository(tempFile)
	result, err := repo.GetAll()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
}

func TestDeleteExpense(t *testing.T) {
	tempFile := setupTempFile(t)
	defer os.Remove(tempFile)

	expenses := []model.Expense{
		{ID: 1, Description: "Coffee", Amount: 5.99, Date: time.Now()},
		{ID: 2, Description: "Lunch", Amount: 12.50, Date: time.Now()},
	}
	setupRepoWithData(t, tempFile, expenses)

	repo := repository.NewExpenseRepository(tempFile)
	err := repo.Delete(1)
	assert.NoError(t, err)

	remaining, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, remaining, 1)
}
