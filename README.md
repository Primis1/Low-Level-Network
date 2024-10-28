# Low-Level-Things

The project divides into several packages, particularly: `protocols`, `logging` and `clientserver`

### Features

- **Low-level HTTP Requests over TCP**: This package allows sending HTTP requests over a TCP connection without using higher-level HTTP libraries.
- **Custom Logging**: Integrated logging for information and error tracking using the `logging` package. 
    - *Basically it was a draft for my further logging, i will update a little in future projects like: blockchain-cooperation and/or news-bot*


***


### `protocols`

- The `protocols` package is designed for handling low-level HTTP requests over TCP connections and logging using a custom logger.
  This package provides tools for making HTTP requests, parsing responses, and managing logging levels. 

#### TCPHttpReq

The `TCPHttpReq` function is used to initiate a basic HTTP request over a TCP connection.

**Example usage:**

```bash
go run main.go -method GET -host example.com -path / -port 8080
```
The command will connect to the specified host on the given port, send an HTTP GET request, and output the response.

#### NewRequest

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


***


### `logging`

- Like I said, it is a reusable logging, that i will use across some of my projects, with updates in each


#### `infoLog/errorLog`

- Variable which initialize two loggers, we do all logging through those two buddies


#### `getCaller()` 

- It used in cutting the path about the file log/error messages are comming from 

#### `application struct` 

- Struct we should declare and fill-up in files where `logger` is used
- It provides two, obvious, methods: Info() and Error()

**Example usage:**

```go
	log := logging.NewLogger(logging.INFO)
	errMsg := logging.NewLogger(logging.ERR)

// usage in project // 
	if err != nil {
		errMsg.Error("error, connection to localHost: %d,  %v \n", *port, err)
	}
	log.Info("connected to %s: will forward stdin \n", conn.RemoteAddr())

// Output //

22/10/21:00 INFO:    connected to address: will forward stdin
```

*Do not forget to set ENV variable of your working directory*


***


### `client`

- Client package, contains _not just_ client, but aloso a server. They put in the same directory only for structural purpose.
- They should be launched in from two separate executables. They connect to host-system port and _host_ and then communicate with each other


### Server (`server.go`)

The `server.go` file sets up a TCP server that listens on a specified port, receives connections, and processes incoming messages in uppercase.

- **Port**: Default is `8080`, can be set using the `-p` flag.
- **Functionality**:
  - Creates a TCP listener.
  - Accepts incoming connections in an infinite loop.
  - Calls `echoUpper.Echo` to process and return messages in uppercase.

**Run the server**:

```bash
go run server.go -p 8080
```

### Client (`client.go`)

- The client.go file initiates a TCP client that connects to the server and interacts via standard input and output.
- Port: Default is 8080, can be set using the -p flag.
    -   Functionality:
        - Connects to the TCP server.
        - Reads from standard input (terminal) and sends input to the server.
        - Logs sent messages and displays server responses.
        - Handles connection and read/write errors using custom logging.

Run the client:

```bash
go run client.go -p 8080
```

***

I've wrote a ReadMe to introduce you the project, so it will never **ever** launched somewhere outside my laptop.
**bUt** it's purpose was in writing things for education. Therefore _you should not clone and OMG run it on your machine(pls no.)_

BTW, writing `readme` is a skill too. 

If you read so far, I'll appreciate the star;)
