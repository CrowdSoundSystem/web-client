integerPattern = "^[0-9]{1,}$";
nIntegerPattern = "^-?[0-9]{1,}$";
floatPattern = "^[0-9]{1,}\.?[0-9]{0,}$";
percentagePattern = "^[1-9][0-9]?$|^100$"
multiplierPattern =  "^0(\.\d+)?|1\.0$"

// This is super hacky, but it's better than copy pasting
// the modal frame work for every fucking setting. I know
// there has to be a better way, but this works for now.
//
// This seems to always be my experience with javascript. Ugh.
formCallback = null;
formElement = null;

$(document).ready(function() {
    $.ajax({
        url:    "admin/setting",
        method: "GET",
        sucess: function(response) {
            $.each(response.settings, function(key, val) {
                // For most things, we can just set it literally.
                // For others, we need to do some stuff first
                switch (key) {
                    case "filterBuffered":
                        if (val) {
                            val = "Enabled"
                        } else {
                            val = "Disabled"
                        }
                        break;
                    case "inactivityThreshold":
                        val = val + " minute(s)";
                        break;
                    case "skipThreshold":
                        val = 100 * val + "%";
                        break;
                }

                $("#" + key).text(val);
            })
        },
        error: function(result) {
            // TODO: Could do a badge, but fuck it, more things to do
            alert("Error retrieving settings.");
        }
    });
});

function createModal(title, form, callback) {
    $("#settings-title").val(title);

    $("#settings-body").empty();
    $("#settings-body").append(form);

    formCallback = callback;
    formElement = form;
    form.show();
}

function createTextForm(label, pattern, initial) {
    var form = $("#form-text").clone();
    form.attr("pattern", pattern);
    form.find(".control-label").text(label);
    form.find(".value").val(initial);
    form.removeClass("has-error");
    return form;
}

function createBooleanForm(label, enabled) {
    var form = $("#form-boolean").clone();
    form.find(".control-label").text(label);
    form.find(".bool").prop("checked", enabled);
    form.removeClass("has-error");
    return form;
}

function postSetting(setting, value) {
    $.ajax({
        url:    "admin/setting",
        method: "POST",
        data:   { key: setting, value: value },
        sucess: function(response) {
            // TODO: could do a badge, but fuck it, more things to do
            alert("Setting saved!");
            window.location.reload();
        },
        error: function(result) {
            // TODO: Could do a badge, but fuck it, more things to do
            alert("Error saving setting: " + result);
            window.location.reload();
        }
    });
}

$("#btn-save").click(function() {
    // Retrieve validation pattern from div.
    var pattern = formElement.attr("pattern");
    var value   = null

    if (formElement.find(".bool").length > 0) {
        value = formElement.find(".bool").is(":checked");
    } else {
        value = formElement.find(".value").val();
    }

    if (pattern) {
        pattern = new RegExp(pattern);
        if (pattern.test(value)) {
            formCallback(value);
        } else {
            $("#form-text").addClass("has-error");
            return;
        }
    } else {
        formCallback(value);
    }

    $("#settingsModal").modal("hide");
    formCallback = null;
    formElement = null;
});

/***************************************
* DB Settings                          *
****************************************/
$("#filterBuffered").click(function() {
    var enabled = $("#filterBuffered").text() == "Enabled";
    createModal("Filter Buffered", createBooleanForm("Enabled", enabled), function(val) {
        postSetting("filterBuffered", val);

        if (val) {
            $("#filterBuffered").text("Enabled");
        } else {
            $("#filterBuffered").text("Disabled");
        }
    });
});

$("#inactivityThreshold").click(function() {
    var initial = $("#inactivityThreshold").text();
    initial = initial.split(" ")[0];

    createModal("Inactivity Threshold", createTextForm("Threshold (minutes)",integerPattern, initial), function(val) {
        postSetting("inactivityThreshold", val * 60 * 1000);
        $("#inactivityThreshold").text(val + " minute(s)");
    });
});

$("#resultLimit").click(function() {
    createModal("Result Limit", createTextForm("Result Limit", nIntegerPattern, $("#resultLimit").text()), function(val) {
        postSetting("resultLimit", val);
        $("#resultLimit").text(val);
    });
});


/***************************************
* Qeueu Settings                       *
****************************************/
$("#session-name").click(function() {
    createModal("Session Name", createTextForm("Session Name", "^[a-zA-Z0-9]{1,20}$", $("#session-name").text()), function(val) {
        postSetting("sessionName", val);
        $("#session-name").text(val);
    });
});

$("#queueSize").click(function() {
    createModal("Queue Size", createTextForm("Queue Size", integerPattern, $("#queueSize").text()), function(val) {
        postSetting("queueSize", queueSize);
        $("#queueSize").text(val);
    });
});

$("#trendingArtistsSize").click(function() {
    createModal("Trending Artists Size", createTextForm("Trending Artists Size", integerPattern, $("#trendingArtistsSize").text()), function(val) {
        postSetting("trendingArtistsSize", val);
        $("#trendingArtistsSize").text(val);
    });
});

$("#skipThreshold").click(function() {
    createModal("Skip Threshold", createTextForm("Skip Threshold [0, 100]", percentagePattern, $("#skipThreshold").text()), function(val) {
        postSetting("skipThreshold", val/100.0);
        $("#skipThreshold").text(val + "%");
    });
});

/***************************************
* Algo Settings                        *
****************************************/
$("#countWeight").click(function() {
    createModal("Count Weight", createTextForm("Count Weight [0.0, 1.0]", multiplierPattern, $("#countWeight").text()), function(val) {
        postSetting("countWeight", val);
        $("#countWeight").text(val);
    });
});

$("#voteWeight").click(function() {
    createModal("Vote Weight", createTextForm("Vote Weight [0.0, 1.0]", multiplierPattern, $("#voteWeight").text()), function(val) {
        postSetting("voteWeight", val);
        $("#voteWeight").text(val);
    });
});

$("#genreWeight").click(function() {
    createModal("Genre Weight", createTextForm("Genre Weight [0.0, 1.0]", multiplierPattern, $("#genreWeight").text()), function(val) {
        postSetting("genreWeight", val);
        $("#genreWeight").text(val);
    });
});

$("#artistWeight").click(function() {
    createModal("Artist Weight", createTextForm("Artist Weight [0.0, 1.0]", multiplierPattern, $("#artistWeight").text()), function(val) {
        postSetting("artistWeight", val);
        $("#artistWeight").text(val);
    });
});

$("#playedAgainMult").click(function() {
    createModal("Played Again Multiplier", createTextForm("Played Again Multiplier [0.0, 1.0]", multiplierPattern, $("#playedAgainMult").text()), function(val) {
        postSetting("playedAgainMultiplier", val);
        $("#playedAgainMult").text(val);
    });
});

$("#minRepeatWindow").click(function() {
    var initial = $("#minRepeatWindow").text();
    initial = initial.split(" ")[0];

    createModal("Minimum Repeat Window", createTextForm("Minimum Repeat Window (Minutes)", integerPattern, initial), function(val) {
        postSetting("minRepeatWindow", val);
        $("#minRepeatWindow").text(val + " minute(s)");
    });
});



