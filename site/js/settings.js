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
    updateVersionInfo();
    updateDBStats();
    updateSettings();
});

function updateVersionInfo() {
    $.ajax({
        url:    "admin/version",
        method: "GET",
        success: function(response) {
            var info = JSON.parse(response);
            $.each(info, function(key, val) {
                $("#" + key).text(val);
            });
        },
        error: function(result) {
            // TODO: Could do a badge, but fuck it, more things to do
            alert("Error retrieving version info.");
        }
    });
}

function updateDBStats() {
    $.ajax({
        url:    "admin/db_stats",
        method: "GET",
        success: function(response) {
            var stats = JSON.parse(response);
            $.each(stats, function(key, val) {
                $("#" + key).text(val);
            });
        },
        error: function(result) {
            // TODO: Could do a badge, but fuck it, more things to do
            alert("Error retrieving version info.");
        }
    });
}

function updateSettings() {
    $.ajax({
        url:    "admin/setting",
        method: "GET",
        success: function(response) {
            var settings = JSON.parse(response);
            $.each(settings, function(key, val) {
                // For most things, we can just set it literally.
                // For others, we need to do some stuff first
                switch (key) {
                    case "filter_buffered":
                        if (val) {
                            val = "Enabled"
                        } else {
                            val = "Disabled"
                        }
                        break;
                    case "inactivity_threshold":
                        val = val / 60 / 1000 + " minute(s)";
                        break;
                    case "skip_threshold":
                        val = 100 * val + "%";
                        break;
                    case "vote_weight":
                        switch (val) {
                            case 0:
                                val = "Low";
                                break;
                            case 1:
                                val = "Equal";
                                break;
                            case 2:
                                val = "High";
                                break;
                        }
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
};


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

function postSetting(setting, type, value) {
    $.ajax({
        url:    "admin/setting",
        method: "POST",
        data:   { key: setting, type: type, value: value },
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
    } else if (formElement.find(":checked").length > 0) {
        value = formElement.find(":checked").val();
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
$("#filter_buffered").click(function() {
    var enabled = $("#filter_buffered").text() == "Enabled";
    createModal("Filter Buffered", createBooleanForm("Enabled", enabled), function(val) {
        postSetting("filter_buffered", "bool", val);

        if (val) {
            $("#filter_buffered").text("Enabled");
        } else {
            $("#filter_buffered").text("Disabled");
        }
    });
});

$("#inactivity_threshold").click(function() {
    var initial = $("#inactivity_threshold").text();
    initial = initial.split(" ")[0];

    createModal("Inactivity Threshold", createTextForm("Threshold (minutes)",integerPattern, initial), function(val) {
        postSetting("inactivity_threshold", "int", val * 60 * 1000);
        $("#inactivity_threshold").text(val + " minute(s)");
    });
});

$("#result_limit").click(function() {
    createModal("Result Limit", createTextForm("Result Limit", nIntegerPattern, $("#result_limit").text()), function(val) {
        postSetting("result_limit", "int", val);
        $("#result_limit").text(val);
    });
});


/***************************************
* Qeueu Settings                       *
****************************************/
$("#session_name").click(function() {
    createModal("Session Name", createTextForm("Session Name", "^[a-zA-Z0-9]{1,20}$", $("#session_name").text()), function(val) {
        postSetting("session_name", "string", val);
        $("#session_name").text(val);
    });
});

$("#queue_size").click(function() {
    createModal("Queue Size", createTextForm("Queue Size", integerPattern, $("#queue_size").text()), function(val) {
        postSetting("queue_size", "int", queue_size);
        $("#queue_size").text(val);
    });
});

$("#trending_artists_size").click(function() {
    createModal("Trending Artists Size", createTextForm("Trending Artists Size", integerPattern, $("#trending_artists_size").text()), function(val) {
        postSetting("trending_artists_size", "int", val);
        $("#trending_artists_size").text(val);
    });
});

$("#skip_threshold").click(function() {
    createModal("Skip Threshold", createTextForm("Skip Threshold [0, 100]", percentagePattern, $("#skip_threshold").text()), function(val) {
        postSetting("skip_threshold", "float", val/100.0);
        $("#skip_threshold").text(val + "%");
    });
});

/***************************************
* Algo Settings                        *
****************************************/
$("#count_weight").click(function() {
    createModal("Count Weight", createTextForm("Count Weight [0.0, 1.0]", multiplierPattern, $("#count_weight").text()), function(val) {
        postSetting("count_weight", "float", val);
        $("#count_weight").text(val);
    });
});

$("#vote_weight").click(function() {
    var form = $("#form-vote-weight").clone();
    switch ($("#vote_weight").text()) {
        case "Low":
            form.find("#optLow").prop("checked", true);
            break;
        case "Equal":
            form.find("#optEqual").prop("checked", true);
            break;
        case "Low":
            form.find("#optHigh").prop("checked", true);
            break;
    }
    createModal("Vote Weight", form, function(val) {
        postSetting("vote_weight", "int", val);
        switch (val) {
            case "0":
                val = "Low";
                break;
            case "1":
                val = "Equal";
                break;
            case "2":
                val = "High";
                break;
        }
        $("#vote_weight").text(val);
    });
});

$("#genre_weight").click(function() {
    createModal("Genre Weight", createTextForm("Genre Weight [0.0, 1.0]", multiplierPattern, $("#genre_weight").text()), function(val) {
        postSetting("genre_weight", "float", val);
        $("#genre_weight").text(val);
    });
});

$("#artist_weight").click(function() {
    createModal("Artist Weight", createTextForm("Artist Weight [0.0, 1.0]", multiplierPattern, $("#artist_weight").text()), function(val) {
        postSetting("artist_weight", "float", val);
        $("#artist_weight").text(val);
    });
});

$("#played_again_mult").click(function() {
    createModal("Played Again Multiplier", createTextForm("Played Again Multiplier [0.0, 1.0]", multiplierPattern, $("#played_again_mult").text()), function(val) {
        postSetting("played_again_multiplier", "float", val);
        $("#played_again_mult").text(val);
    });
});

$("#min_repeat_window").click(function() {
    var initial = $("#min_repeat_window").text();
    initial = initial.split(" ")[0];

    createModal("Minimum Repeat Window", createTextForm("Minimum Repeat Window (Minutes)", integerPattern, initial), function(val) {
        postSetting("min_repeat_window", "int", val);
        $("#min_repeat_window").text(val + " minute(s)");
    });
});



