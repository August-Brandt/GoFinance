package reader

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type Transaction struct {
	Amount      float32 `csv:"Beløb"`
	Description string  `csv:"Tekst"`
	Type        string  `csv:"Hovedkategori"`
	Subtype     string  `csv:"Kategori"`
	Comment string `csv:"Kommentar"`
	Balance float32 `csv:"Saldo"`
	Date string `csv:"Dato"`
}

func check(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}

func ReadTransactions(filePath string) []Transaction {
	// TODO: Find a way for the reader to correctly parse floats
	// with formats like: x.xxx,xx

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

	return transactions
}
