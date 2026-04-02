# Fund Daily Performance Alert

A Go-based AWS Lambda application that monitors stock fund performance and sends daily email alerts via Amazon SNS.

## Overview

This application fetches daily stock quotes from Alpha Vantage API, processes fund data stored in Amazon DynamoDB, and sends performance summary emails to subscribers through Amazon SNS.

## Architecture

```
┌─────────────────┐     ┌──────────────────┐     ┌─────────────────┐
│   Alpha Vantage │────▶│   AWS Lambda     │────▶│  Amazon SNS     │
│   (Stock Quotes)│     │   (Go Runtime)   │     │  (Email Alert)  │
└─────────────────┘     └──────────────────┘     └─────────────────┘
                              │
                              ▼
                        ┌──────────────────┐
                        │  Amazon DynamoDB │
                        │  (Fund Storage)  │
                        └──────────────────┘
```

## Features

- **Automated Stock Quote Fetching**: Retrieves real-time stock quotes from Alpha Vantage API
- **Fund Management**: Stores and retrieves fund-stock mappings from DynamoDB
- **Performance Calculation**: Calculates daily percentage changes for each stock
- **Email Notifications**: Sends formatted performance reports via Amazon SNS
- **API Rate Limiting**: Built-in rate limiting to comply with Alpha Vantage API limits

## Project Structure

```
fund-alerts/
├── src/
│   ├── main.go                 # Lambda entry point
│   ├── Config.go               # AWS and API configuration handling
│   ├── FundsRepository.go      # DynamoDB repository interface
│   ├── FundsService.go         # Business logic for fund processing
│   ├── QuotesFetcher.go        # Alpha Vantage API integration
│   ├── EmailTextHandler.go     # Email content generation
│   ├── SNSHandler.go           # SNS notification handler
│   ├── SecurityHelper.go       # KMS decryption utilities
│   └── resources/
│       └── configs.yml         # Application configuration
├── scripts/
│   ├── local-build.sh          # Local build script
│   └── compress-main-file.sh   # Deployment package compression
├── go.mod                      # Go module definition
├── go.sum                      # Go dependencies checksum
└── README.md                   # This file
```

## Prerequisites

- Go 1.13 or higher
- AWS CLI configured with appropriate credentials
- AWS Lambda execution role with permissions for:
  - DynamoDB (Scan, Query)
  - SNS (Publish)
  - KMS (Decrypt) - if using encrypted configuration values
- Alpha Vantage API key

## Configuration

Edit `src/resources/configs.yml` with your settings:

```yaml
aws:
  region: us-east-1
  dynamodb:
    table: FundStock
  sns:
    arn: <your-sns-topic-arn>

alphavantage:
  endpoint: https://www.alphavantage.co
  apikey: <your-api-key>
```

### DynamoDB Table Schema

The `FundStock` table should contain items with the following structure:

| Attribute | Type | Description |
|-----------|------|-------------|
| FundStockKey | String | Composite key: `{FUND_CODE}-{SYMBOL}` |
| Company | String | Company name |
| Symbol | String | Stock ticker symbol |
| FundName | String | Human-readable fund name |

## Building

### Local Build

Run the local build script to create a deployment package:

```bash
cd scripts
./local-build.sh
```

This will:
1. Compile the Go code for Linux (amd64)
2. Copy the resources directory
3. Create a `main.zip` deployment package

### Manual Build

```bash
GOARCH=amd64 GOOS=linux go build -o bin/main .
```

## Deployment

1. Build the deployment package:
   ```bash
   cd scripts && ./local-build.sh
   ```

2. Deploy to AWS Lambda:
   ```bash
   aws lambda update-function-code \
     --function-name your-function-name \
     --zip-file fileb://bin/main.zip
   ```

3. Configure a CloudWatch Events rule to trigger the function on your desired schedule (e.g., daily at market close).

## Usage

The Lambda function is triggered automatically based on your CloudWatch Events schedule. When invoked, it:

1. Loads configuration from `configs.yml`
2. Decrypts sensitive values (API keys, SNS ARN) using AWS KMS
3. Scans the DynamoDB table for all fund-stock mappings
4. Groups stocks by fund
5. Fetches current quotes from Alpha Vantage API
6. Generates a performance report email
7. Publishes the report to the configured SNS topic

## Dependencies

- [github.com/aws/aws-lambda-go](https://github.com/aws/aws-lambda-go) - AWS Lambda Go runtime
- [github.com/aws/aws-sdk-go](https://github.com/aws/aws-sdk-go) - AWS SDK for Go
- [gopkg.in/yaml.v2](https://gopkg.in/yaml.v2) - YAML parser

## License

This project is licensed under the MIT License.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

