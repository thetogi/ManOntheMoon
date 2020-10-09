$( document ).ajaxStart(function() {
    $("body").addClass("busy");
});
$( document ).ajaxComplete (function() {
    $("body").removeClass("busy");
});

function EnvToggle() {

    var e = document.getElementById('envParams');
    if (typeof (e) != 'undefined' && e != null) {
        e.parentNode.removeChild(e);
    } else {
        $.ajax({
            url: "http://localhost:8080/toggleShowEnv",
            type: 'GET',
            beforeSend: function () {
                // $("#loader").css("display", "block");
                console.log('wait 3 seconds')

                setTimeout(function () {
                    console.log('wait 3 seconds complete');
                }, 3000);
            },

            success: function (data) {
                console.log("EnvToggle Success");
                //alert(data);
                Success = true;
            },
            error: function (textStatus, errorThrown) {
                console.log("EnvToggle Failed");
                console.log(errorThrown);
                //alert(errorThrown);
                Success = false;
            }
        }).done(function (data) {
            $('#env').append(data);
          //  $("body").removeClass("busy");
            // $("#loader").css("display", "none");
        });
    }
}

function TestLoader() {
    $("body").addClass("busy");
    setTimeout(function () {
        $("body").removeClass("busy");
        console.log('wait 3 seconds complete');
    }, 3000);

}