function doRequest(
    method = 'POST',
    credentials = 'same-origin',
    mode = 'no-cors',
    apiUrl,
    bodyType = 'text',
    headers = {},
) {

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
    fetch(apiUrl, initData)
        .then(response => {
            return response.text();
        })
        .then(data => {
            if (data.trim() === "") {
                alert("JS code doesn't have access to the response data. You can check the server response in tools like BurpSuite");
            } else {
                alert("Request successfully executed. Response data: " + data.trim());
            }
            console.log("Response data:", data);
        })
        .catch(function (error) {
            alert("An error is occurred: " + error + "\nPlease check the console log for more details");
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
        const apiDomain = formData.get("server");
        const header = formData.get("header")
        const headerValue = formData.get("headerValue");
        const bodyType = formData.get("bodyType");
        const isHttp = formData.get("isHttp");

        const headers = {}
        if (header && headerValue) {
            headers[header] = headerValue
        }
        let apiUrl = "https://" + apiDomain
        if (isHttp) {
            apiUrl = "http://" + apiDomain
        }
        doRequest(method, credentials, mode, apiUrl, bodyType, headers);
    });
});

