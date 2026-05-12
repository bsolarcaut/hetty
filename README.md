# Hetty

Hetty is an HTTP toolkit for security research. It aims to become an open source
alternative to commercial software like Burp Suite Pro, with powerful features
tailored to the needs of the infosec and bug bounty community.

![Screenshot](docs/screenshot.png)

## Features

- Man-in-the-middle (MITM) HTTP/1.1 proxy, with TLS support via generated CA
- Project based database storage (SQLite), for isolated and portable environments
- Scope support, to help keep work focused
- Intercept requests and responses for manual review/editing
- Search and filter through proxy logs

## Requirements

- Go 1.21+
- Node.js 18+ (for frontend development)

## Installation

### From source

```bash
git clone https://github.com/dstotijn/hetty.git
cd hetty
go build -o hetty ./cmd/hetty
```

### Docker

```bash
docker pull dstotijn/hetty
docker run -v $HOME/.hetty:/root/.hetty -p 8080:8080 dstotijn/hetty
```

## Usage

```
Usage of hetty:
  -addr string
        TCP address to listen on, in the form "host:port" (default ":8080")
  -adminPath string
        File path to admin build directory (default "$HOME/.hetty/admin")
  -cert string
        CA certificate filepath. Creates a new CA certificate if file doesn't exist (default "$HOME/.hetty/hetty_cert.pem")
  -key string
        CA private key filepath. Creates a new CA private key if file doesn't exist (default "$HOME/.hetty/hetty_key.pem")
  -db string
        Database file path (default "$HOME/.hetty/hetty.db")
  -verbose
        Enable verbose logging
```

### Setting up the CA certificate

Hetty generates a CA certificate and private key on first run. To intercept
TLS traffic, you need to install the CA certificate in your browser or OS.

The CA certificate is stored at `$HOME/.hetty/hetty_cert.pem` by default.

#### macOS

To trust the certificate on macOS, run:

```bash
sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain $HOME/.hetty/hetty_cert.pem
```

After adding it, you may need to restart your browser for the changes to take effect.

## Development

### Backend

```bash
go run ./cmd/hetty
```

### Frontend

```bash
cd admin
npm install
npm run dev
```

## Contributing

Contributions are welcome! Please read the [contributing guidelines](CONTRIBUTING.md)
before submitting a pull request.

## License

[Apache License 2.0](LICENSE)
