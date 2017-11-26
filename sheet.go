package main

import (
	sheets "google.golang.org/api/sheets/v4"
)

type Spreadsheet struct {
	*sheets.Spreadsheet
	transactions *TransactionSheet
}

func (ss *Spreadsheet) SetTransactionSheet() {
	for _, s := range ss.Sheets {
		if s.Properties.Title == "Transactions" {
			ss.transactions = &TransactionSheet{Sheet: s}
			ss.transactions.FindColumns()
		}
	}
}

type TransactionSheet struct {
	*sheets.Sheet
	DescriptionColumn int
	CategoryColumn    int
	AmountColumn      int
}

func (s *TransactionSheet) FindColumns() {
	for _, d := range s.Data {
		for i, v := range d.RowData[0].Values {
			switch v.FormattedValue {
			case "Description":
				s.DescriptionColumn = i
			case "Category":
				s.CategoryColumn = i
			case "Amount":
				s.AmountColumn = i
			}
		}
	}
}

func (s *TransactionSheet) Transactions() []Transaction {
	t := []Transaction{}
	for _, r := range s.Data[0].RowData {
		t = append(t, Transaction{
			Description: r.Values[s.DescriptionColumn],
			Category:    r.Values[s.CategoryColumn],
			Amount:      r.Values[s.AmountColumn],
		})
	}

	return t
}

type Transaction struct {
	Description *sheets.CellData
	Category    *sheets.CellData
	Amount      *sheets.CellData
}
