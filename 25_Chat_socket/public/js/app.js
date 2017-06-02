$(document).ready(function() {

    var user_name;

    $("#form_registro").on("submit", function(e) {
        e.preventDefault();
        user_name = $("#user_name").val();

        $.ajax({
            type: "POST",
            url: "http://localhost:8000/validate",
            data: {
                "user_name": user_name
            },
            success: function(data) {
                result(data)
            }
        })
    })

    function result(data) {
        obj = JSON.parse(data);
        if (obj.isvalid) {
            create_conexion()
        }else{
            console.log('Intentalo con otro nombre de usuario');
        }
    }

    function create_conexion() {
        var conexion = WebSocket("ws://loclhost:8000/chat" + user_name);
    }

});
