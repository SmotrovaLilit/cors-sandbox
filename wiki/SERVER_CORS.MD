## Server CORS implementation

- Server allows requests from origin https://origin.cors.com, http://origin.cors.com .
- Server allows methods GET, POST, PUT, and HEAD.
- Server allows headers Content-Type, Custom-Allowed-Header.
- Server allows sending credentials

```go
    handler := handlers.CORS(
handlers.AllowedOrigins([]string{"http://127.0.0.1:3000"}),
handlers.AllowedHeaders([]string{"Content-Type"}),
handlers.AllowCredentials(),
handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"}),
)(routerAnotherServer)
```

## References
1. https://developer.mozilla.org/ru/docs/Web/HTTP/CORS
