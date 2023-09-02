let apiCrossOriginUrl = 'http://127.0.0.1:8081/api';
let apiSameOriginUrl = 'http://127.0.0.1:3000/api';

function doRequest(
    method = 'POST',
    credentials = 'same-origin',
    mode = 'no-cors',
    sameOrigin = false,
    bodyType= 'text',
    headers = {},
) {

    let url = apiCrossOriginUrl
    if (sameOrigin) {
        url = apiSameOriginUrl
    }
    initData = {
        method: method,
        headers: headers,
        credentials: credentials,
        mode: mode
    }
    if (method === 'POST' || method === 'PUT') {
        initData.body = "Hello World!"
        if (bodyType === 'json') {
            initData.headers['Content-Type'] = 'application/json'
            initData.body = JSON.stringify({message: "Hello World!"})
        }
    }
    console.log("Request data:", initData);
    fetch(url, initData)
        .then(response => response.text())
        .then(data => {
            console.log("Response data:", data);
        })
        .catch(function (error) {
            console.log("Error:", error);
        });

}

document.addEventListener("DOMContentLoaded", function () {
    const form = document.getElementById("cors-form");

    form.addEventListener("submit", function (event) {
        event.preventDefault();

        const formData = new FormData(form);
        const mode = formData.get("mode");
        const credentials = formData.get("credentials");
        const method = formData.get("method");
        const sameOrigin = formData.get("server") === "same";
        const header = formData.get("header")
        const headerValue = formData.get("headerValue");
        const bodyType = formData.get("bodyType");

        const headers = {}
        if (header && headerValue) {
            headers[header] = headerValue
        }

        doRequest(method, credentials, mode, sameOrigin, bodyType, headers);
    });
});

