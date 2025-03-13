package main

import (
	"expense-tracker-app/internal/repository"
	"expense-tracker-app/internal/service"
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'add', 'list', 'summary' or 'delete' subcommand")
		return
	}

	repo := repository.NewExpenseRepository("storage/expenses.json")
	svc := service.NewExpenseService(repo)

	switch os.Args[1] {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		description := addCmd.String("description", "", "Expense description")
		amount := addCmd.Float64("amount", 0, "Expense amount")

		_ = addCmd.Parse(os.Args[2:])
		if *description == "" || *amount <= 0 {
			fmt.Println("Invalid input. Description and amount are required.")
			os.Exit(1)
		}

		id, err := svc.AddExpense(*description, *amount)
		if err != nil {
			fmt.Println("Failed to add expense:", err)
			os.Exit(1)
		}

		fmt.Printf("Expense added successfully (ID: %d \n", id)
	case "list":
		expenses, err := svc.GetAllExpenses()
		if err != nil {
			fmt.Println("Failed to get expenses:", err)
			os.Exit(1)
		}

		fmt.Println("ID  Date        Description  Amount")
		for _, exp := range expenses {
			fmt.Printf("%d  %s  %s  $%.2f\n", exp.ID, exp.Date.Format("2006-01-02"), exp.Description, exp.Amount)
		}

	case "summary":
		summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
		month := summaryCmd.Int("month", 0, "Month for summary (optional)")
		_ = summaryCmd.Parse(os.Args[2:])

		total, err := svc.Summary(*month)
		if err != nil {

		}

		if *month == 0 {
			fmt.Printf("Total expenses: $%.2f\n", total)
		} else {
			fmt.Printf("Total expenses for month %d: $%.2f\n", *month, total)
		}

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCmd.Int("id", 0, "ID of expense to delete")

		_ = deleteCmd.Parse(os.Args[2:])
		if *id == 0 {
			fmt.Println("Invalid input. Expense ID is required.")
			os.Exit(1)
		}

		err := svc.DeleteExpense(*id)
		if err != nil {
			fmt.Println("Failed to delete expense:", err)
			os.Exit(1)
		}

		fmt.Println("Expense deleted successfully.")

	default:
		fmt.Println("Invalid command. Use 'add', 'list', 'summary', or 'delete'.")
	}
}
