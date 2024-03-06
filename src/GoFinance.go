package main

import (
	"fmt"
	"flag"
	"GoFinance/src/Reader"
)

func main() {
	file := flag.String("file", "unset", "Path to the csv file with transaction history")
	flag.Parse()
	
	if *file == "unset" {
		panic("The file was not set")
	}

	fmt.Println("File is: " + *file)
	
	transactionHistory := Reader.ReadTransactions(*file)
	showTransactions(transactionHistory)
}

func showTransactions(th Reader.TransactionHistory) {
	for i, transaction := range th.Transactions {
		fmt.Printf("Transaction %d\n\tAmount: %f\n\tType: %s\n\tSubtype: %s\n\tDescription: %s\n", 
			i, transaction.Amount, transaction.Type, transaction.Subtype, transaction.Description)
	} 
}

