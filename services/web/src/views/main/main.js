import React from 'react';
import './main.css';

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

export default function main({
                                 channels, currChannel, chooseChannel,
                                 activeUserUUID, users,
                                 messages, sendMessage,
                                 currMessageValue, setCurrMessageValue,
                             }) {
    let activeUser = users[activeUserUUID] || {};
    return (
        <div className="mainBody d-flex flex-column">
            <header className="d-flex flex-row justify-content-between">
                <img src="/logo512.png" className="navbar-brand mx-2 py-2" alt="" />
                <img className="py-2 px-3 d-block"
                     src={window.APIHost + "/api/avatars/" + activeUser.uuid}
                     alt={activeUser.name}/>
            </header>
            <div className="d-flex flex-row flex-grow-1">
                <div className="sidebar py-3 d-flex flex-column">
                    <div className="channels-list">
                        <div className="d-flex flex-row justify-content-between mx-3">
                            <div>
                                <i className="feather icon-chevron-down mr-2"/> Channels
                            </div>
                            <button className="btn">
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
                                <div className="messages-list d-flex m-3 flex-grow-0 flex-shrink-0">
                                    <div className="flex-grow-0 flex-shrink-0">
                                        {((messages && messages[currChannel.uuid]) || []).map(message => <div className="d-flex flex-column">
                                            <b>{(users[message.user_uuid] || {}).name} {new Date(message.time).toLocaleTimeString()}</b>

                                            {message.text}
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