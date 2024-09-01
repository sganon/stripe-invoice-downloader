# stripe-invoice-downloader

A command-line tool to download Stripe invoices within a specified date range.

## Usage

```bash
stripe-invoice-downloader [global options]
```

### Global Options

- `--api-key`: Your Stripe API key (required)
- `--from`: Start date of the export in UTC (format: YYYY-MM-DD HH:MM)
- `--to`: End date of the export in UTC (format: YYYY-MM-DD HH:MM)
- `--out-dir`: The output directory for downloaded invoices (default: "out")

## Environment Variables
- `STRIPE_API_KEY`: Can be used instead of the `--api-key` flag


## Build

To build the project, follow these steps:

1. Ensure you have Go installed on your system. You can download it from [golang.org](https://golang.org/dl/).

2. Clone the repository:
```bash
   git clone https://github.com/yourusername/stripe-invoice-downloader.git
   cd stripe-invoice-downloader
```

3. Build the project:
```bash
   go build -o stripe-invoice-downloader
```

This will create an executable named `stripe-invoice-downloader` in your current directory.

4. (Optional) To install the tool globally, run:
```bash
   go install
```

Now you can use the `stripe-invoice-downloader` command from anywhere on your system.

