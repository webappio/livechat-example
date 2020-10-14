export default class ProtocolHandler {
    ws
    dashboard

    constructor(dashboard) {
        this.dashboard = dashboard;
    }

    send(message) {
        this.ws.send(JSON.stringify(message));
    }

    processMessage({type, ...data}) {
        switch(type) {
            case "redirect-to-login":
                this.dashboard.setState({currPage: "/login"})
                return
            case "user-info":
                this.dashboard.setState({activeUser: data["user"]})
                return
            case "channels":
                this.dashboard.setState(prevState => {
                    return {
                        channels: data["channels"],
                        activeChannelUUID: prevState.activeChannelUUID || data["channels"][0].uuid,
                    }
                })
                return
            default:
                console.warn("unexpected message: "+type);
        }
    }

    init() {
        this.ws = new WebSocket(
            (document.location.protocol === "http:" ? "ws://" : "wss://")
            + window.APIHost
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