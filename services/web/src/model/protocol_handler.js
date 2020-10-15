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
                this.dashboard.setState({activeUserUUID: data["uuid"]})
                return
            case "users":
                let newUsers = {};
                for(let user of data["users"]) {
                    newUsers[user.uuid] = user;
                }
                this.dashboard.setState({users: newUsers})
                return
            case "channels":
                this.dashboard.setState(prevState => {
                    return {
                        channels: data["channels"],
                        activeChannelUUID: prevState.activeChannelUUID || data["channels"][0].uuid,
                    }
                })
                return
            case "channel_messages":
                this.dashboard.setState(prevState => {
                    let newMessages = {...(prevState.messages || {})};
                    for(let message of data["messages"]) {
                        newMessages[message.channel_uuid] = [...(newMessages[message.channel_uuid] || []), message]
                    }
                    for(let key of Object.keys(newMessages)) {
                        newMessages[key].sort((a, b) => a.index-b.index)
                        let newArray = [];
                        for(let message of newMessages[key]) {
                            if(newArray.length === 0 || newArray[newArray.length-1].index !== message.index) {
                                newArray.push(message);
                            }
                        }
                        newMessages[key] = newArray;
                    }
                    return {messages: newMessages}
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