package main

import (
	"fmt"
	"flag"
	"GoFinance/src/Reader"
)

func main() {
	file := flag.String("file", "", "Path to the csv file with transaction history")
	flag.Parse()
	
	fmt.Println("File is: " + *file)
	
	// test reader
	if *file != "" {
		transactionHistory := Reader.ReadTransactions(*file)
		for i, transaction := range transactionHistory.Transactions {
			fmt.Printf("Transaction %d\n\tAmount: %f\n\tType: %s\n\tSubtype: %s\n\tDescription: %s\n", 
				i, transaction.Amount, transaction.Type, transaction.Subtype, transaction.Description)
		}
		
	}
}