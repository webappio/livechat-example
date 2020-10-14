import React from 'react';
import './main.css';


export default function main({channels, currChannel, activeUser}) {
    return (
        <div>
            <header className="d-flex flex-row justify-content-end">
                <img src={"/avatars/" + activeUser.uuid} alt="User Avatar"/>
            </header>
            <div className="d-flex flex-column">
                <div className="sidebar d-flex flex-column">
                    <div className="channels-list">
                        <i className="feather icon-chevron-down"/> Channels
                        <button><i className="feather icon-plus"/></button>
                        {channels.map(channel => <button key={channel.name} className="ml-3">
                            # {channel.name}
                        </button>)}
                    </div>
                    <div className="dms-list">
                        <i className="feather icon-chevron-down"/> Direct messages
                        <button><i className="feather icon-plus"/></button>
                    </div>
                </div>
                <div className="messages d-flex flex-column">

                </div>
            </div>
        </div>
    );
}