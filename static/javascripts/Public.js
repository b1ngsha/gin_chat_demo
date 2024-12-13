function WebSocketConnect(user_info, to_user_info = null) {
    if ("WebSocket" in window) {
        console.log("您的浏览器支持 WebSocket!");

        let ws = new WebSocket("ws://127.0.0.1:8000/ws/");
        console.log(ws)
    }
}