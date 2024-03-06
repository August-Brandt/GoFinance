package Reader

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type TransactionHistory struct {
	Transactions []Transaction
}

type Transaction struct {
	Amount      float32 `csv:"amount"`
	Description string  `csv:"description"`
	Type        string  `csv:"type"`
	Subtype     string  `csv:"subtype"`
}

func check(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}

func ReadTransactions(filePath string) TransactionHistory {
	// Set up reader
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return r
	})

	// Open file
	csvFile, err := os.Open(filePath)
	check(err)
	defer csvFile.Close()
	// Read file
	var transactions []Transaction
	err = gocsv.UnmarshalFile(csvFile, &transactions)
	check(err)

	return TransactionHistory{
		Transactions: transactions,
	}
}
