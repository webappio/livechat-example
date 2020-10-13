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
            channels: {}
        };
    }

    componentDidMount() {
        this.protoHandler.init();
    }

    render() {
        return <Router>
            <Switch>
                <Route path="/">
                    <Main/>
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
