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

class Index extends React.Component {
    constructor(props) {
        super(props);

        this.protoHandler = new ProtocolHandler(this);
        this.state = {
            currPage: "/",
            channels: [],
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
                        activeUser={this.state.activeUser}
                        channels={this.state.channels}
                        currChannel={this.state.channels.filter(({uuid}) => uuid === this.state.activeChannelUUID).pop()}
                        sendMessage={() => this.protoHandler.send({"type": "new_message", "contents": this.state.currMessage})}
                        currMessageValue={this.state.currMessage}
                        setCurrMessageValue={newValue => this.setState({currMessage: newValue})}
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
