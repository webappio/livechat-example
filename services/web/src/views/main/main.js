import React from 'react';
import './main.css';

function TextArea(props) {
    return <div className="message-area-wrapper m-3">
        <textarea className="message-area"
                  value={props.value}
                  onKeyDown={e => {
                      if (e.key === "Enter" && e.ctrlKey) {
                          props.setValue(props.value+"\n");
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
                                 channels, currChannel, activeUser,
                                 sendMessage, currMessageValue, setCurrMessageValue,
                             }) {
    return (
        <div className="mainBody d-flex flex-column">
            <header className="d-flex flex-row justify-content-end">
                <img className="p-2 d-block" src={"/avatars/" + activeUser.uuid} alt="User Avatar"/>
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
                        {channels.map(channel => <button key={channel.name} className="ml-3">
                            # {channel.name}
                        </button>)}
                    </div>
                </div>
                <div className="messages d-flex flex-column justify-content-between">
                    <div className="messages-list">

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