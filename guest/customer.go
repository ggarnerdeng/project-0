package guest

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

// Customer data
type Customer struct {
	userName string
	password string
	name     string
	balance  float64
}

// NewCustomer is a Constructor for Customer
func NewCustomer(userName string, password string,
	name string, balance float64) *Customer {
	n := Customer{
		userName: userName,
		password: password,
		name:     name,
		balance:  balance,
	}

	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	db.Exec("INSERT INTO customer"+"(userName,password,name,balance)"+
		"VALUES($1,$2,$3,$4)", userName, password, name, balance)

	return &n
}

func (a Customer) String() string {
	var output string
	t := fmt.Sprintf("%.2f", a.balance)
	output = a.userName + " | " + a.password + " | " + a.name + " | $" + t + "\n"
	return output
}

// Balance returns the amount of money in a customer's balance
func (a *Customer) Balance() float64 {
	return a.balance
}

// Withdraw removes money from a customer's balance
func (a *Customer) Withdraw(i float64) {
	if a.balance < i {
		fmt.Println("Insufficient funds, transaction canceled")
	} else {
		a.balance -= i
	}

}

// Deposit adds money to a customer's balance
func (a *Customer) Deposit(i float64) {
	a.balance += i
}

// Transfer moves money from one customer to another customer's balance
func (a *Customer) Transfer(i float64, b *Customer) {
	if a.balance < i {
		fmt.Println("Insufficient funds, transaction canceled")
	} else {
		a.Withdraw(i)
		b.balance += i
	}
}