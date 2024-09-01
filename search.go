package main

import (
	"fmt"
	"time"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/invoice"
)

func searchInvoices(start, end time.Time) ([]*stripe.Invoice, error) {
	params := &stripe.InvoiceSearchParams{
		SearchParams: stripe.SearchParams{Query: fmt.Sprintf("created>%d AND created<%d AND status:\"%s\"", start.Unix(), end.Unix(), stripe.InvoiceStatusPaid)},
	}
	result := invoice.Search(params)
	invoices := []*stripe.Invoice{}

	for result.Next() {
		inv := result.Invoice()
		invoices = append(invoices, inv)
	}

	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("result error: %w", err)
	}

	return invoices, nil
}
