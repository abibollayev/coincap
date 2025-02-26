# CoinCap

CoinCap is a simple CLI tool that fetches real-time cryptocurrency prices using WebSockets from CoinMarketCap.

## Features
- Live price updates for BTC and ETH
- Colored output with arrows indicating price changes
- Uses WebSockets for real-time data

## Prerequisites
Before running the project, ensure you have the following installed:
- Go (latest version)

## Installation & Setup

### 1. Clone the repository:
```sh
git clone https://github.com/abibollayev/coincap.git
cd coincap
```

### 2. Install dependencies
```sh
go mod tidy
```

### 2. Run the application:
```sh
go run main.go btc
```

###    or for Ethereum::
```sh
go run main.go eth
```

## Example Output
```sh
BTC: 84298.30$ ↑
```
The arrow (↑ or ↓) indicates whether the price has increased or decreased.

## Stopping the Application
Press CTRL+C to stop the CLI tool.

## Future Improvements
- Support for more cryptocurrencies
- WebSocket reconnect on failure
- Display additional market data like volume & market cap
