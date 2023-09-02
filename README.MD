# CORS
Using this sandbox you can test CORS requests with options:
- mode
- credentials
- headers
- methods
- request content type

You can send request to the same origin or to cross-origin.
![example.png](assets/example.png)
## Client site settings Mode
### Mode cors (default) 
If request is not same-origin:
- For methods GET, HEAD, POST, the browser will send the request. JS can read the response only if the server returns the right CORS headers.
- For other methods, the browser firstly send OPTIONS method to the server to ask permissions. If the server returns the right CORS headers, the browser will send the request. JS can read the response only if the server returns the right CORS headers.
- In some case for GET, HEAD, POST the browser will send options too. For example, if you use custom headers, or send application/json content-type. Read more in [3](#references)

### Mode no-cors
If request is not same-origin:
- Request method is limited to GET, HEAD or POST. The browser will **do** request to the server but js will **not** have access to the response.
- If request method is not one of these the browser will return an error and the request will **not** be sent to the server.

### Mode same-origin.
Use when you want to send request to the same origin as the calling script. 
If you use this mode and the origin isn't the same, you will get an error.
The request will **not** be sent to the server.

## Server CORS implementation
- Server allow requests from origin http://127.0.0.1:3000 with credentials. 
- Server allow methods GET, POST, PUT, HEAD.
- Server allow headers Content-Type.
```go
    handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://127.0.0.1:3000"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"}),
	)(routerAnotherServer)
```
## References
1. https://developer.mozilla.org/en-US/docs/Web/API/Request/mode
2. https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch
3. https://developer.mozilla.org/ru/docs/Web/HTTP/CORS