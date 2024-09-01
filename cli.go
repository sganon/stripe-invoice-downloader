package main

import (
	"fmt"
	"time"

	"github.com/stripe/stripe-go/v79"
	"github.com/urfave/cli/v2"
)

func initCliApp() *cli.App {
	app := &cli.App{
		Name: "stripe-invoice-downloader",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "api-key", Usage: "Your Stripe API key", EnvVars: []string{"STRIPE_API_KEY"}},
			&cli.StringFlag{Name: "from", Usage: "Start date of the export in UTC, format 2006-01-02 15:04"},
			&cli.StringFlag{Name: "to", Usage: "End date of the export in UTC, format 2006-01-02 15:04"},
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

			invoices, err := searchInvoices(*from, *to)
			if err != nil {
				return fmt.Errorf("error searching invoices: %w", err)
			}

			if err := downloadInvoices(invoices, cCtx.String("out-dir")); err != nil {
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
