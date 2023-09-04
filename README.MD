# CORS

The sandbox features:
1. Availability to test cross-origin requests using fetch[1](#references) with different settings:
    - mode
    - credentials
    - headers
    - methods
    - content type 
2. Availability to test behaviour with sending cookie in cross-origin requests with different cookie settings:
    - SameSite mode (None, Lax, Strict, Default)
    - HttpOnly flag
    - Secure flag
3. Availability to test cross-origin requests and sending cookie behaviour with sending requests to:
    - subdomains/different domains
    - http/https
    - server with allowed cross-origin requests/server with not allowed cross-origin requests 

![example.png](assets/example.png)


The result can be seen in the console log.
Using tools like BurpSuite you can see real requests to the server with responses.

## Installation

Add to etc hosts

```bash
127.0.0.1 allowing.cors.com
127.0.0.1 inallowing.cors.com
127.0.0.1 allowing.cors-another.com
127.0.0.1 inallowing.cors-another.com
127.0.0.1 origin.cors.com
```

Add certs.caroot.cer to trusted certificates

Run docker-compose

```bash
docker-compose up
```

Run http servers
```bash
go run cmd/server/main.go
```

Open application

Go to http://origin.cors.com or https://origin.cors.com

## Run the sandbox in unsafe browser mode
If you need to test behavior without browser security policy you can run chrome in unsafe mode.

Close all chrome instances and run in terminal:

MacOS

```bash
open -na Google\ Chrome --args --user-data-dir=/tmp/temporary-chrome-profile-dir --disable-web-security --disable-site-isolation-trials
```

See details here [2](#references)

Kali Linux

```bash
google-chrome  --user-data-dir="~/google-chrome-data" --disable-web-security
```

Go to http://origin.cors.com or https://origin.cors.com

## Update certificates
If you need to update certificates you can:

```bash
brew install mkcert
mkcert -install
```

move created certificate to nginx/certs folder

```bash
cd nginx/certs
mkcert "*.cors.com"
mkcert "*.cors-another.com"
```
## Notes about CORS behaviour
[INDEX.MD](wiki%2FINDEX.MD)

## References
1. https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch
2. https://stackoverflow.com/questions/35432749/disable-web-security-in-chrome-48
3. https://developer.mozilla.org/ru/docs/Web/HTTP/CORS
