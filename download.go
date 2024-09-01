package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/stripe/stripe-go/v79"
)

func ensureOutputDir(outputDir string) error {
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to check output directory: %w", err)
	}
	return nil
}

func downloadInvoices(invoices []*stripe.Invoice, outputDir string) error {
	if err := ensureOutputDir(outputDir); err != nil {
		return err
	}

	for _, invoice := range invoices {
		if invoice.InvoicePDF == "" {
			continue
		}

		fileName := fmt.Sprintf("invoice_%s.pdf", invoice.Number)
		filePath := filepath.Join(outputDir, fileName)

		resp, err := http.Get(invoice.InvoicePDF)
		if err != nil {
			return fmt.Errorf("failed to download invoice %s: %w", invoice.ID, err)
		}
		defer resp.Body.Close()

		out, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", filePath, err)
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return fmt.Errorf("failed to write invoice %s to file: %w", invoice.ID, err)
		}

		fmt.Printf("Downloaded invoice %s to %s\n", invoice.ID, filePath)
	}

	return nil
}
