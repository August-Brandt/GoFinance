package main

import (
	"fmt"
	"flag"
	"GoFinance/reader"
)

func main() {
	file := flag.String("file", "unset", "Path to the csv file with transaction history")
	flag.Parse()
	
	if *file == "unset" {
		panic("The file was not set")
	}

	fmt.Println("File is: " + *file)
	
	transactionHistory := reader.ReadTransactions(*file)
	showTransactions(transactionHistory)
}

func showTransactions(th []reader.Transaction) {
	for i, transaction := range th {
		fmt.Printf("Transaction %d\n" + 
					"\tAmount: %f\n" + 
					"\tDescription: %s\n" + 
					"\tType: %s\n" + 
					"\tSubtype: %s\n" + 
					"\tComment: %s\n" + 
					"\tBalance: %f\n" + 
					"\tDate: %s\n", 
					i, 
					transaction.Amount, 
					transaction.Description, 
					transaction.Type, 
					transaction.Subtype, 
					transaction.Comment,
					transaction.Balance,
					transaction.Date,
				)
	} 
}

func createDataSet(th reader.Transaction) {
	// TODO: Create dataset map from the transaction history
}