//
// app.js
//
// This file contains the logic of backend communication
// to keep all the data up to date.
//

$(document).read(function() {
    if (!("Websocket" in window)) {
        // TODO: Nicer transition.
        alert("Websockets not supported.");
    } else {
        // Setup conection and callback handlers.
    }
})

function onOpen() {
    // TODO: Enable the screen, or something
}

function onMessage(message) {
    // TODO: Determine events and act accordingly
}

function onError(err) {
    // TODO: Close the socket. This is important for chrome
}

function onClose() {
    // TODO: Disable the screen, or something
}
