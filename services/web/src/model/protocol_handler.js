export default class ProtocolHandler {
    ws
    dashboard

    constructor(dashboard) {
        this.dashboard = dashboard;
    }

    send(message) {
        this.ws.send(message);
    }

    processMessage({type, ...data}) {
        console.log("got message: " + type);
        switch(type) {
            case "redirect-to-login":
                this.dashboard.setState({currPage: "/login"})
                return
            default:
                console.warn("unexpected message: "+type);
        }
    }

    init() {
        this.ws = new WebSocket(
            (document.location.protocol === "http:" ? "ws://" : "wss://")
            + (document.location.host === "localhost:3000" ? "localhost:8000" : document.location.host)
            + "/api/ws"
        )
        this.ws.binaryType = "blob";

        this.ws.addEventListener("close", () => {
            setTimeout(() => this.init(), 1000);
        });
        this.ws.addEventListener("message", msg => {
            try {
                this.processMessage(JSON.parse(msg.data));
            } catch (err) {
                console.error(err);
            }
        });
    }
}