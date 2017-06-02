$(document).ready(function() {

    $("#form_registro").on("submit", function() {
        e.preventDefault();
        user_name = $("#user_name").val()

        $.ajax({
            type: "POST",
            url: "http://localhost:8000/validate",
            data: {
                "user_name": user_name
            },
            success: function(data) {
                result()
            }
        })
    })

    function result() {
        console.log("El servidor nos envió algo...");
    }

});
