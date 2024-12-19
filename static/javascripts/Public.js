const msgTypeOnline = 1;
const msgTypeGetUserList = 4;

var ws;

/** scroll to lowest **/
function toLow () {
    $('.scrollbar-masosx.scroll-content.scroll-scrolly_visible').animate({
        scrollTop: $('.scrollbar-macosx.scroll-content.scroll-scrolly_visible').prop('scrollHeight')
    }, 500);
}

function formatTime (time) {
    let date = new Date(time + 8 * 3600 * 1000);
    return date.toJSON().substring(0, 19).replace('T', ' ');
}

function WebSocketConnect (user_info) {
    if ("WebSocket" in window) {
        console.log("Your browser supports WebSocket!");

        ws = new WebSocket("ws://127.0.0.1:8080/ws");

        let chat_info = $('.main .chat_info');

        ws.onopen = function () {
            console.log("Connected to websocket server");
            // render chat_info div
            chat_info.html(chat_info.html() + '<li class="system_info"><span> ✅ websocket connection has established </span></li>');
            // send message
            let message = JSON.stringify({
                "status": msgTypeOnline,
                "from_user_id": user_info.user_id.toString(),
                "room_id": user_info.room_id,
                "avatar_id": user_info.avatar_id,
                "username": user_info.username,
            });
            console.log("Send message to server: " + message);
            ws.send(message);
            toLow();
        }

        ws.onmessage = function (event) {
            let received_msg = JSON.parse(event.data);
            let time = formatTime(received_msg.time);
            console.log("Received:", received_msg);
            switch (received_msg.status) {
                case msgTypeOnline:
                    chat_info.html(chat_info.html() + '<li class="system_info"> <span>' + "[" + received_msg.username + "] " + time + " enter" + '</span></li>');
                    break;
                case msgTypeGetUserList:
                    $(".popover-title").html(received_msg.count + "users online.")
                    $.each(received_msg.list, function (index, item) {
                        if (received_msg.user_id == item.user_id) {
                            $('.ul-user-list').html(
                                $('.ul-user-list').html() +
                                '<li style="pointer-events: none;" class="li-user-item" data-user_id=' + item.user_id + ' data-username=' + item.username + ' data-room_id=' + item.room_id + ' data-avatar_id=' + item.avatar_id + '  ><img src="/static/images/user/' + item.avatar_id + '.png" alt=""><b>' + " " + item.username + '</b>' + '</li>'
                            );
                        } else {
                            $('.ul-user-list').html(
                                $('.ul-user-list').html() +
                                '<li class="li-user-item" data-user_id=' + item.user_id + ' data-username=' + item.username + ' data-room_id=' + item.room_id + ' data-avatar_id=' + item.avatar_id + '><img src="/static/images/user/' +
                                item.avatar_id + '.png" alt=""><b>' + " " + item.username + '</b>' + '</li>'
                            );
                        }
                    })
            }
        }

        ws.onclose = function (event) {
            console.log("WebSocket connection closed");
        }

        ws.onerror = function (error) {
            console.log("WebSocket error: ", error);
        }
    }
}