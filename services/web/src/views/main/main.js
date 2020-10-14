import React from 'react';
import './main.css';

function TextArea(props) {
    return <div className="message-area-wrapper m-3">
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
                                 channels, currChannel, chooseChannel, activeUser,
                                 sendMessage, currMessageValue, setCurrMessageValue,
                             }) {
    return (
        <div className="mainBody d-flex flex-column">
            <header className="d-flex flex-row justify-content-end">
                <img className="py-2 px-3 d-block" src={window.APIHost + "/api/avatars/" + activeUser.uuid}
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
                    <div className="messages-list">
                        <div className="channel-info-bar m-3">
                            {!currChannel ? null : (
                                <b>
                                    <i className="feather icon-lock"/> {currChannel.name}
                                </b>
                            )}
                        </div>
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