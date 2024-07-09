package reader

import (
	"os"
	"path/filepath"
	"testing"
)

func createCsvFile(csvString string) error {
	err := os.WriteFile("testfile.csv", []byte(csvString), 0644)
	if err != nil {
		return err
	}
	return nil
}

func cleanUp() {
	os.Remove("testfile.csv")
}

func TestReadTransaction(t *testing.T) {
	// Setup
	csvString := `Dato;Tekst;Beløb;Saldo;Hovedkategori;Kategori;Kommentar
10.06.2024;Test Transaction 1;-150,00;600,00;Mad og indkøb;Dagligvarer;""
10.06.2024;Test Transaction 2;5,00;658,00;Indtægter;Anden indtægt;""
06.06.2024;Test Transaction 3;-108,50;700,00;Transport;Taxa og offentlig transport;""
04.06.2024;Test Transaction 4;-30,00;234,00;Fritid;Café, restaurant og bar;""`

	err := createCsvFile(csvString)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanUp()

	// Run function that gets tested
	transactions := ReadTransactions("testfile.csv")

	// Assert
	if len(transactions) != 4 {
		t.Errorf("Incorrect number of transactions\nExpected: 4\nGot: %d", len(transactions))
	}
	if transactions[0].Amount != -150.00 {
		t.Errorf("Incorrect number for amount\nExpected; -150.00\nGot: %f", transactions[0].Amount)
	}
	if transactions[1].Balance != 658.00 {
		t.Errorf("Incorrect number for balance\nExpected; 658.00\nGot: %f", transactions[1].Balance)
	}
}

func TestNormalizeCsvFile(t *testing.T) {
	err := normalizeCsvFile("../../data/data.csv")
	if err != nil {
		t.Fatalf("Error in normalizeCsvFile: %v", err)
	}
}

func TestReadTransactionWithNormalization(t *testing.T) {
	// Setup
	csvString := `Dato;Tekst;Beløb;Saldo;Hovedkategori;Kategori;Kommentar
10.06.2024;Test Transaction 1;-150,00;600,00;Mad og indkøb;Dagligvarer;""
10.06.2024;Test Transaction 2;5,00;658,00;Indtægter;Anden indtægt;""
06.06.2024;Test Transaction 3;-1.008,50;700,00;Transport;Taxa og offentlig transport;""
04.06.2024;Test Transaction 4;-30,00;1.234,00;Fritid;Café, restaurant og bar;""`

	err := createCsvFile(csvString)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanUp()

	// Run function that gets tested
	normalizeCsvFile("testfile.csv")
	transactions := ReadTransactions(filepath.Join(os.TempDir(), "temp.csv"))

	// Assert
	if len(transactions) != 4 {
		t.Errorf("Incorrect number of transactions\nExpected: 4\nGot: %d", len(transactions))
	}
	if transactions[2].Amount != -1008.50 {
		t.Errorf("Incorrect number for amount\nExpected; -1008.50\nGot: %f", transactions[0].Amount)
	}
	if transactions[3].Balance != 1234.00 {
		t.Errorf("Incorrect number for balance\nExpected; 1234.00\nGot: %f", transactions[1].Balance)
	}
}
