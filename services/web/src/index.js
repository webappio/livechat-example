import React from 'react';
import {
    BrowserRouter as Router,
    Switch,
    Route
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
            channels: [],
            activeUser: {},
            activeChannelUUID: "",
        };
    }

    componentDidMount() {
        this.protoHandler.init();
    }

    render() {
        return <Router>
            <Switch>
                <Route path="/">
                    <Main
                        activeUser={this.state.activeUser}
                        channels={this.state.channels}
                        currChannel={this.state.channels.filter(({uuid}) => uuid === this.state.activeChannelUUID).pop()}
                    />
                </Route>
                <Route path="/login">
                    <Login/>
                </Route>
            </Switch>
        </Router>
    }
}

ReactDOM.render(
    <React.StrictMode><Index/></React.StrictMode>,
    document.getElementById('root')
);
