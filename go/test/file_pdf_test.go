package test

import (
	"testing"

	"github.com/ThiagoRodriguesdeSantana/desafio_conductor/go/file"
	"github.com/ThiagoRodriguesdeSantana/desafio_conductor/go/model"
)

func TestFile(t *testing.T) {

	t.Run("test file pdf", func(t *testing.T) {

		transactions := make([]model.Transaction, 0)

		transactions = append(transactions, model.Transaction{
			AccountId:   "ASSdDasdfadfa-asdfafa-asdfa",
			Description: "Netflix",
			Value:       10.00,
		})
		transactions = append(transactions, model.Transaction{
			AccountId:   "ASSdDasdfadfa-asdfafa-asdfa",
			Description: "Amazon",
			Value:       20.00,
		})
		transactions = append(transactions, model.Transaction{
			AccountId:   "ASSdDasdfadfa-asdfafa-asdfa",
			Description: "Netflix",
			Value:       100.00,
		})

		file := file.NewPDF()

		path := file.GeneratePDF(transactions)

		if len(path) == 0 {
			t.Fail()
		}
	})
}
