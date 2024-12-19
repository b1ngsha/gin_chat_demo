$(document).ready(function() {
    $("#chat-list-li-top").click(function() {
        let offset = $("#hidden-chat-list-li-top").attr("data-offset");
        let room_id = $(".room").attr("data-room_id");
        let user_id = getURLParams("user_id");

        $.ajax({
            url: `/pagination?room_id=${room_id}&offset=${offset}&user_id=${user_id}`,
            success: function(data) {
                
            }
        })
    })
})

function getURLParams(key) {
    return decodeURIComponent((new RegExp('[?|&]' + key + '=' + '([^&;]+?)(&|#|;|$)', "ig").exec(location.search) || [, ""])[1].replace(/\+/g, '%20')) || '';
}