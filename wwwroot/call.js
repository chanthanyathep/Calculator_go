function displayResult() {
    var input = document.getElementById("screen").innerHTML;
    var Input = {
        Input: input,
    }
    var requestPayload = JSON.stringify(Input);

    const xhttp = new XMLHttpRequest();
    xhttp.onload = function() {
        var statusCode = this.status;
        var responseBody = this.responseText;
        if (statusCode === 200) {
            document.getElementById("output").value = responseBody;
        }
    }
    xhttp.open("POST", "/getValue");
    xhttp.setRequestHeader('Content-Type', 'application/json');
    xhttp.send(requestPayload);
}