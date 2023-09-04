# Client side settings

## Mode

Fetching another origin

|                                                                                                                                                   | cors                                                                                                                                                                               | no-cors                                                              | same-origin                                                          |
|---------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------|----------------------------------------------------------------------|
| GET, HEAD, POST                                                                                                                                   | Outgoing request will be sent \| JS can read response if the server sends right CORS headers                                                                                       | Outgoing request will be sent \| JS can't read response              | Outgoing request will not be sent \|The browser will return an error |
| PUT, DELETE, PATCH                                                                                                                                | Preflight request[4](#references) \| Outgoing request will be sent, If the server allows to do CORS\| JS can read the response only if the server returns the right CORS headers.  | Outgoing request will not be sent \|The browser will return an error | Outgoing request will not be sent \|The browser will return an error |
| some cases with POST, GET, HEAD. For example, if custom headers are used, or content type is been application/json. Read more in [3](#references) | Preflight request[4](#references) \| Outgoing request will be sent, If the server allows to do CORS \| JS can read the response only if the server returns the right CORS headers. | Outgoing request will be sent \| JS can't read response              | Outgoing request will not be sent \|The browser will return an error | 

## Credentials

Credentials are Cookies and HTTP Authentication

## References

1. https://developer.mozilla.org/en-US/docs/Web/API/Request/mode
2. https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch
3. https://developer.mozilla.org/ru/docs/Web/HTTP/CORS
4. https://developer.mozilla.org/en-US/docs/Glossary/Preflight_request