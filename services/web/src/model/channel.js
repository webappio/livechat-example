export class Channel {
    constructor(name) {
        this.name = name;
        this.members = [];
        this.messages = [];
    }

    addMessage(message) {
        this.messages = [...this.messages, message];
    }
}