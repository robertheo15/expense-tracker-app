# Expense Tracker
https://roadmap.sh/projects/expense-tracker

Build a simple expense tracker application to manage your finances. The application should allow users to 
add, delete, and view their expenses. The application should also provide a summary of the expenses.

## How to run

Clone the repository and run the following command:

```bash
git clone https://github.com/robertheo15/expense-tracker-app.git
cd github-activity-app
```


### Run the following command run the project:
```
$ expense-tracker add --description "Lunch" --amount 20
# Expense added successfully (ID: 1)

$ expense-tracker add --description "Dinner" --amount 10
# Expense added successfully (ID: 2)

$ expense-tracker list
# ID  Date       Description  Amount
# 1   2024-08-06  Lunch        $20
# 2   2024-08-06  Dinner       $10

$ expense-tracker summary
# Total expenses: $30

$ expense-tracker delete --id 2
# Expense deleted successfully

$ expense-tracker summary
# Total expenses: $20

$ expense-tracker summary --month 8
# Total expenses for August: $20
```

### Run the following command run the test case project:
```
go test -v ./internal/repository                                                                                                                                                                                            ─╯
go test -v ./internal/service                                                                                                                                                                                            ─╯
```