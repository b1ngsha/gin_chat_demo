$(document).ready(function() {
    $("#chat-list-li-top").click(function() {
        let offset = $("#hidden-chat-list-li-top").attr("data-offset");
        let room_id = $(".room").attr("data-room_id");
        let user_id = getURLParams("user_id");

        $.ajax({
            url: `/pagination?room_id=${room_id}&offset=${offset}&user_id=${user_id}`,
            success: function(data) {
                let item = data.data.list;
                if (item == null) {
                    layer.msg("no more messsage!");
                    $("#hidden-chat-list-li-top").attr("data-offset", offset);
                    $("#chat-list-li-top").hide();
                    return false;
                }

                $.each(item, function(index, value) {
                    if (value.from_user_id == $("#body-room").attr("data-user_id")) {
                        $("#chat-list-li-top").after(
                            `<li class="right">
                                <img src="/static/images/user/${value.avatar_id}.png" alt="">
                                <b>${value.username}</b>
                                <i>${value.created_at}</i>
                                <div class="aaa">${value.content}</div>
                            </li>`
                        );
                    } else {
                        $("#chat-list-li-top").after(
                            `<li class="left">
                                <img src="/static/images/user/${value.avatar_id}.png" alt="">
                                <b>${value.username}</b>
                                <i>${value.created_at}</i>
                                <div class="aaa">${value.content}</div>
                            </li>`
                        );
                    }
                })
                $("hidden-chat-list-li-top").attr("data-offset", parseInt(offset) + 100);
            }
        })
    })
})

function getURLParams(key) {
    return decodeURIComponent((new RegExp('[?|&]' + key + '=' + '([^&;]+?)(&|#|;|$)', "ig").exec(location.search) || [, ""])[1].replace(/\+/g, '%20')) || '';
}