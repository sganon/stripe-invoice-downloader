package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/paymentmethod"
	"github.com/urfave/cli/v2"
)

func initCliApp() *cli.App {
	app := &cli.App{
		Name: "stripe-invoice-downloader",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "api-key", Usage: "Your Stripe API key", EnvVars: []string{"STRIPE_API_KEY"}},
			&cli.StringFlag{Name: "from", Usage: "Start date of the export in UTC, format 2006-01-02 15:04"},
			&cli.StringFlag{Name: "to", Usage: "End date of the export in UTC, format 2006-01-02 15:04"},
			&cli.StringFlag{Name: "expand", Usage: "Fields to expand in the invoice search, comma-separated"},
			&cli.StringFlag{Name: "out-dir", Usage: "The output directory where the downloaded invoices are stored", Value: "out"},
		},
		Action: func(cCtx *cli.Context) error {
			apiKey := cCtx.String("api-key")
			stripe.Key = apiKey
			from, err := timeFlagValue(cCtx.String("from"))
			if err != nil {
				return fmt.Errorf("error parsing from: %w", err)
			}
			to, err := timeFlagValue(cCtx.String("to"))
			if err != nil {
				return fmt.Errorf("error parsing to: %w", err)
			}

			expand := parseExpandFlag(cCtx.String("expand"))

			invoices, err := searchInvoices(*from, *to, expand)
			if err != nil {
				return fmt.Errorf("error searching invoices: %w", err)
			}

			var sepaInvoice []*stripe.Invoice
			for _, inv := range invoices {
				if inv.Charge != nil {
					pm, err := paymentmethod.Get(inv.Charge.PaymentMethod, nil)
					if err != nil {
						fmt.Printf("Error fetching PaymentMethod: %v\n", err)
						continue
					}
					fmt.Printf("Invoice ID: %s, PaymentMethod: %s\n", inv.ID, pm.Type)
					if pm.Type == "sepa_debit" {
						sepaInvoice = append(sepaInvoice, inv)
					}
				} else {
					fmt.Printf("Invoice ID: %s, No Charge associated\n", inv.ID)
				}
			}

			if err := downloadInvoices(sepaInvoice, cCtx.String("out-dir")); err != nil {
				return fmt.Errorf("error downloading invoices: %w", err)
			}

			return nil
		},
	}

	return app
}

const timeFlagLayout = "2006-01-02 15:04"

func timeFlagValue(v string) (*time.Time, error) {
	t, err := time.Parse(timeFlagLayout, v)
	if err != nil {
		return nil, fmt.Errorf("error parsing time value: %w", err)
	}

	return &t, nil
}

func parseExpandFlag(expand string) []*string {
	if expand == "" {
		return nil
	}
	fields := strings.Split(expand, ",")
	expandFields := make([]*string, len(fields))
	for i, field := range fields {
		expandFields[i] = &field
	}
	return expandFields
}
