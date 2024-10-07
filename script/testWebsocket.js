fetch ("http://localhost:5050/posts", {
    method: "GET",
    headers: {
        "Content-Type": "application/json"
    }   
    }).then(function(response) {
        return response.json();
    }).then(function(json){
        console.log(json)
    })

var ws = new WebSocket("ws://localhost:5050/ws");

ws.onopen = function() {
    console.log("Opened");
    ws.send("Hello World");
}

ws.onmessage = function(e) {
    console.log(e.data);
}

ws.onerror = function(e) {
    console.log(e);
}