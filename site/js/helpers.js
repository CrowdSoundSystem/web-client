//
// helpers.js
//
// This file contains helpers to update various
// components of the UI.
//

function updateSessionInfo(info) {
    $("#stat-name").text(info.sessionName);
    $("#stat-users").text(info.users);
}

function updateSkipStatus(voteCount, totalUsers) {
    // This works. Javascript!
    var percentage = 0.0;
    if (totalUsers > 0) {
        percentage = (voteCount / totalUsers) * 100;
    }

    $("#skip-status-progress").width(percentage + '%')
}

function updateNowPlaying(track) {
    $("#np-title").text(track.name);
    $("#np-artist").text(track.artist);
    $("#np-genre").text(track.genre);
}

function updateTopArtists(artists) {
    // Clear existing items
    $("#trending-artists").find(".list-group-item").remove();

    var max = Math.min(artists.length, 5)
    for (var i = 0; i < max; i++) {
        $("#trending-artists").append(
                $("<li>").text(artists[i]).attr("class", "list-group-item")
        );
    }
}

function updateSongList(list, tracks) {
    list.find(".list-group-item").remove();

    var template = $("<li>").attr("class", "list-group-item");
    for (var i = 0; i < tracks.length; i++) {
        var item = template.clone();

        var name = "";
        if (tracks[i].artist) {
            name = tracks[i].artist + " - "
        }
        name += tracks[i].name;

        item.text(name);
        item.append($("<span>").attr("class", "badge").text(tracks.votes));
        list.append(item);
    }
}

function updateUpNext(tracks) {
    updateSongList($("#buffered-queue"), tracks)
}

function updateDynamicQueue(tracks) {
    updateSongList($("#dynamic-queue"), tracks)
}
