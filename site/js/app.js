//
// app.js
//
// This file contains the logic of backend communication
// to keep all the data up to date.
//

$(document).ready(function() {
    if (!window.WebSocket) {
        // TODO: Nicer transition.
        alert("Websockets not supported.");
    } else {
        // Setup conection and callback handlers.
        connection = new WebSocket("ws://" + window.location.host + "/event_stream");
        connection.onopen = onOpen
        connection.onmessage = onMessage
        connection.onerror = onError
        connection.onclose = onClose
    }
})

function onOpen() {
    // TODO: Enable the screen, or something
    console.log("Websocket connection open")
}

function onMessage(message) {
    var eventData = JSON.parse(message.data);
    if (eventData.error) {
        console.log("unexpected stream error: " + eventData.error);
        return
    }

    switch (eventData.eventType) {
        case "queue":
            updateUpNext(eventData.event.buffered);
            updateDynamicQueue(eventData.event.queued);
            break;
        case "now_playing":
            updateNowPlaying(eventData.event.song);
            break;
        case "session_data":
            updateSessionInfo(eventData.event);
            break;
        case "trending_artists":
            updateTopArtists(eventData.event.artists);
            break;
        case "skip_status":
            updateSkipStatus(eventData.event.vote_to_skip, eventData.event.total_users);
    }
}

function onError(err) {
    // TODO: Close the socket. This is important for chrome
    connection.onClose()
}

function onClose() {
    // TODO: Disable the screen, or something
    console.log("Websocket connection closed")
}

$("#btn-force-skip").click(function() {
    $("#btn-force-skip").addClass("disabled");
    $.ajax({
        url:    "admin/skip",
        error: function(xhr, status, error) {
            alert("Could not skip: " + error);
        },
        complete: function(xhr, status) {
            $("#btn-force-skip").removeClass("disabled");
        }
    });
});
