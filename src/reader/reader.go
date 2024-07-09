package reader

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocarina/gocsv"
)

type Transaction struct {
	Amount      float32 `csv:"Beløb"`
	Description string  `csv:"Tekst"`
	Type        string  `csv:"Hovedkategori"`
	Subtype     string  `csv:"Kategori"`
	Comment     string  `csv:"Kommentar"`
	Balance     float32 `csv:"Saldo"`
	Date        string  `csv:"Dato"`
}

func check(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}

func normalizeCsvFile(path string) error {
	// Change floating point numbers in csv file to be
	// formatted like: xxxx.xx instead of x.xxx,xx

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var buf bytes.Buffer

	scanner := bufio.NewScanner(file)

	// Added csv header to buffer
	scanner.Scan()
	buf.Write(scanner.Bytes())
	buf.Write([]byte("\n"))

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ";")
		// Normalize amount
		s := strings.Replace(splitLine[2], ".", "", -1)
		splitLine[2] = strings.Replace(s, ",", ".", -1)

		// Normalize balance
		s = strings.Replace(splitLine[3], ".", "", -1)
		splitLine[3] = strings.Replace(s, ",", ".", -1)

		// Write normalized line to buffer
		newLine := strings.Join(splitLine, ";")
		buf.Write([]byte(newLine))
		buf.Write([]byte("\n"))
	}

	// Create normalized file in temp directory
	err = os.WriteFile(filepath.Join(os.TempDir(), "temp.csv"), buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadTransactions(filePath string) []Transaction {
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
