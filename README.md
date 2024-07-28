# Go HTTP Server from Scratch

This Go project implements a basic HTTP server from scratch using TCP sockets. It does not use the `net/http` package and is designed to handle HTTP requests and responses at a low level.

## Features

- **Custom HTTP Server**: Implemented using raw TCP connections.
- **Configurable Port**: You can specify the port on which the server listens, with a default value of 6969.
- **Request and Response Handling**: Parses and formats HTTP requests and responses according to the HTTP specification.

## Running the Server

To run the HTTP server, use the following command:

```sh
go run main.go -port <port>
```
If no port is specified, the server will default to port 6969.

## HTTP Request Format

The server expects incoming HTTP requests in the following format:
```
Method Request-URI HTTP-Version CRLF
headers CRLF
message-body
```
- Method: The HTTP method (e.g., GET, POST).
- Request-URI: The requested resource's URI.
- HTTP-Version: The HTTP version (e.g., HTTP/1.1).
- CRLF: Carriage return and line feed (\r\n).
- headers: Request headers.
- message-body: Optional body of the request.
## HTTP Response Format


The server sends HTTP responses in the following format:

```
HTTP-Version Status-Code Reason-Phrase CRLF
headers CRLF
message-body
```

- HTTP-Version: The HTTP version (e.g., HTTP/1.1).
- Status-Code: The status code indicating the result of the request (e.g., 200, 404).
- Reason-Phrase: A textual description of the status code (e.g., OK, Not Found).
- CRLF: Carriage return and line feed (\r\n).
- headers: Response headers.
- message-body: Optional body of the response.

This project is licensed under the MIT License - see the LICENSE file for details.
