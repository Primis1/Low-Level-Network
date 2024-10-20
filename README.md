# Low-Level-Things

The projet divides into several different packages, particularly: `protocols`, `logging` and `clientserver`

### Features

- **Low-level HTTP Requests over TCP**: This package allows sending HTTP requests over a TCP connection without using higher-level HTTP libraries.
- **Custom Logging**: Integrated logging for information and error tracking using the `logging` package. 
    - *Basically it was a draft for my further logging,
      which you will see in projects like: `news-bot` and `ethical-ddos`*
- **Request/Response Structs**: Convenient structs for creating HTTP requests and parsing responses.

### `protocols`

- The `protocols` package is designed for handling low-level HTTP requests over TCP connections and logging using a custom logger.
  This package provides tools for making HTTP requests, parsing responses, and managing logging levels. 

#### TCPHttpReq Function

The `TCPHttpReq` function is used to initiate a basic HTTP request over a TCP connection.

**Example usage:**

```bash
go run main.go -method GET -host example.com -path / -port 8080
```

The command will connect to the specified host on the given port, send an HTTP GET request, and output the response.

#### NewRequest Function

Use `NewRequest` to create a new HTTP request programmatically.

**Example:**

```go
req := protocols.NewRequest("localhost", "/", "GET", "")
```

- `host`: The target host.
- `url`: The URL path (must begin with `/`).
- `method`: HTTP method (e.g., GET, POST).
- `body`: Optional body for methods like POST.

### Command-line Arguments

- `-method`: HTTP method (default: `GET`).
- `-host`: The host for the resource (default: `localhost`).
- `-path`: The URL path (default: `/`).
- `-port`: The port to connect to (default: `8080`).

