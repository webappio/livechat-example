import React from 'react';
import './main.css';

class NewChannelLightbox extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            channelName: "",
            channelDescription: "",
        }
    }

    render() {
        return <div className="new-channel-lightbox"
                    style={{display: this.props.visible ? "flex" : "none"}}>
            <div className="new-channel-modal p-3">
                <div className="d-flex flex-row justify-content-between">
                    <h5>Create a channel</h5>
                    <button onClick={() => this.props.onClose()} className="btn">
                        <i className="feather icon-x"/>
                    </button>
                </div>
                <form onSubmit={e => {
                    e.preventDefault();
                    this.props.onChannelCreate({name: this.state.channelName, description: this.state.channelDescription});
                    this.setState({channelName: "", channelDescription: ""});
                    this.props.onClose();
                }}>
                    <div className="form-group">
                        <label htmlFor="new-channel-name">Name</label>
                        <input type="text"
                               className="form-control"
                               id="new-channel-name"
                               name="name"
                               placeholder="general"
                               value={this.state.channelName}
                               onChange={e => this.setState({channelName: e.target.value})}
                        />
                    </div>
                    <div className="form-group">
                        <label htmlFor="new-channel-name">Description</label>
                        <input type="text"
                               className="form-control"
                               id="new-channel-description"
                               name="description"
                               value={this.state.channelDescription}
                               onChange={e => this.setState({channelDescription: e.target.value})}
                        />
                    </div>
                    <div className="form-group">
                        <button className="btn btn-primary">Create</button>
                    </div>
                </form>
            </div>
        </div>
    }
}

function TextArea(props) {
    return <div className="message-area-wrapper m-2 p-2">
        <textarea className="message-area"
                  value={props.value}
                  onKeyDown={e => {
                      if (e.key === "Enter" && e.ctrlKey) {
                          props.setValue(props.value + "\n");
                      }
                  }}
                  onKeyPress={e => {
                      if (e.key === "Enter" && !e.ctrlKey) {
                          props.sendMessage();
                          e.preventDefault();
                          return false;
                      }
                  }}
                  onChange={e => props.setValue(e.target.value)}
        />
    </div>
}

export default class main extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            createChannelModelOpen: false,
        }
    }

    render() {
        let {
            channels, currChannel, chooseChannel,
            activeUserUUID, users,
            messages, sendMessage,
            currMessageValue, setCurrMessageValue,
            onChannelCreate,
        } = this.props;
        let activeUser = users[activeUserUUID] || {};
        return (
            <div className="mainBody d-flex flex-column">
                <header className="d-flex flex-row justify-content-between">
                    <div>
                        <img src="/logo512.png" className="navbar-brand mx-2 py-2" alt=""/>
                        livechat-example
                    </div>
                    <img className="py-2 px-3 d-block"
                         src={window.APIHost + "/api/avatars/" + activeUser.uuid}
                         alt={activeUser.name}/>
                </header>
                <NewChannelLightbox
                    visible={this.state.createChannelModelOpen}
                    onClose={() => this.setState({createChannelModelOpen: false})}
                    onChannelCreate={onChannelCreate}
                />
                <div className="d-flex flex-row flex-grow-1">
                    <div className="sidebar py-3 d-flex flex-column">
                        <div className="channels-list d-flex flex-column align-items-start">
                            <div className="d-flex flex-row justify-content-between px-3 mb-3 w-100">
                                <div>
                                    <i className="feather icon-chevron-down mr-2"/> Channels
                                </div>
                                <button className="btn" onClick={() => this.setState({createChannelModelOpen: true})}>
                                    <i className="feather icon-plus"/>
                                </button>
                            </div>
                            {channels.map(channel => <button
                                key={channel.uuid}
                                className={"btn ml-3" + (currChannel.uuid === channel.uuid ? " active" : "")}
                                onClick={() => chooseChannel(channel.uuid)}
                            >
                                <span className="mr-1">#</span> {channel.name}
                            </button>)}
                        </div>
                    </div>
                    <div className="messages d-flex flex-column justify-content-between">
                        <div>
                            {!currChannel ? null : (
                                <div>
                                    <div className="channel-info-bar m-3">
                                        <b>
                                            <i className="feather icon-lock"/> {currChannel.name}
                                        </b>
                                    </div>
                                    <div className="messages-list-container d-flex m-3 flex-grow-0 flex-shrink-0">
                                        <div className="messages-list flex-grow-0 flex-shrink-0">
                                            {((messages && messages[currChannel.uuid]) || []).map(message => <div
                                                className="d-flex flex-column message-wrapper">
                                                <b className="message-author">{(users[message.user_uuid] || {}).name} {new Date(message.time).toLocaleTimeString()}</b>
                                                <div className="message-text">{message.text}</div>
                                            </div>)}
                                        </div>
                                    </div>
                                </div>
                            )}
                        </div>
                        <TextArea
                            sendMessage={sendMessage}
                            value={currMessageValue}
                            setValue={setCurrMessageValue}
                        />
                    </div>
                </div>
            </div>
        );
    }
}