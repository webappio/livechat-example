import React from 'react';
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Redirect,
} from "react-router-dom";
import ReactDOM from 'react-dom';
import './common.css';
import Main from './views/main/main';
import Login from "./views/login/login";
import ProtocolHandler from "./model/protocol_handler";

window.APIHost = (document.location.host === "localhost:3000" ? "localhost:8000" : document.location.host)

class Index extends React.Component {
    constructor(props) {
        super(props);

        this.protoHandler = new ProtocolHandler(this);
        this.state = {
            currPage: "/",
            channels: [],
            users: {},
            activeUser: {},
            activeChannelUUID: "",
            currMessage: "", //the content of the (unsent) message being written
        };
    }

    componentDidMount() {
        this.protoHandler.init();
    }

    render() {
        return <Router>
            <Redirect to={this.state.currPage} />
            <Switch>
                <Route path="/login">
                    <Login/>
                </Route>
                <Route path="/">
                    <Main
                        activeUserUUID={this.state.activeUserUUID}
                        users={this.state.users}
                        channels={this.state.channels}
                        messages={this.state.messages}
                        currChannel={this.state.channels.filter(({uuid}) => uuid === this.state.activeChannelUUID).pop()}
                        chooseChannel={uuid => this.setState({activeChannelUUID: uuid})}
                        sendMessage={() => {
                            this.protoHandler.send({
                                "type": "new_message",
                                "contents": this.state.currMessage,
                                "channel_uuid": this.state.activeChannelUUID,
                            })
                            this.setState({currMessage: ""})
                        }}
                        currMessageValue={this.state.currMessage}
                        setCurrMessageValue={newValue => this.setState({currMessage: newValue})}
                        onChannelCreate={({name, description}) => {
                            this.protoHandler.send({
                                "type": "new_channel",
                                "name": name,
                                "description": description,
                            })
                        }}
                    />
                </Route>
            </Switch>
        </Router>
    }
}

ReactDOM.render(
    <React.StrictMode><Index/></React.StrictMode>,
    document.getElementById('root')
);
