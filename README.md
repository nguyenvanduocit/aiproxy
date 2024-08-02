# Anthropic API Proxy

This project is a simple proxy server for the Anthropic API, built using Go and the Fiber web framework. It allows you to route requests to the Anthropic API through your own server, which can be useful for adding custom logic, handling CORS issues, or masking the origin of requests.

## Features

- Proxies all requests to the Anthropic API
- Handles CORS (Cross-Origin Resource Sharing) headers
- Configurable port
- Timeout handling for long-running requests

## Prerequisites

- Go 1.16 or later
- Git

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/anthropic-api-proxy.git
   cd anthropic-api-proxy
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

## Usage

1. Set the `PORT` environment variable (optional, defaults to 3000):
   ```
   export PORT=8080
   ```

2. Run the server:
   ```
   go run main.go
   ```

3. The server will start on the specified port (or 3000 if not specified). You can now send requests to `http://localhost:PORT/anthropic/*`, and they will be proxied to the Anthropic API.

## Configuration

- `PORT`: The port on which the server will listen. Defaults to 3000 if not set.
- `DefaultTimeout`: The read and write timeout for the server. Currently set to 5 minutes.

## API

The server proxies all requests to the Anthropic API. Simply replace the Anthropic API base URL with your proxy server's URL:

```
Original: https://api.anthropic.com/v1/complete
Proxied:  http://localhost:PORT/anthropic/v1/complete
```

All request methods (GET, POST, PUT, DELETE, OPTIONS) are supported.

## CORS

CORS is enabled for all origins (`*`) by default. You can modify the CORS settings in the `main.go` file if needed.

## Error Handling

If there's an error proxying the request, the server will return an appropriate error response.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT License](LICENSE)

## Disclaimer

This is an unofficial proxy for the Anthropic API and is not affiliated with or endorsed by Anthropic. Use at your own risk and ensure you comply with Anthropic's terms of service when using their API.
