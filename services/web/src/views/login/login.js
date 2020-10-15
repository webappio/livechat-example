import React from "react";
import "./login.css";

export default function Login(props) {
    return <div className="mainBody">
        <div className="nav ml-3 my-1 d-flex align-items-center">
            <img src="/logo512.png" className="navbar-brand" alt="" />
            Livechat Example
        </div>
        <div className="d-flex flex-column align-items-center justify-content-center w-100 h-auto">
            <div>
                <h1>Login to the Livechat Example</h1>
                <form action={"//"+window.APIHost+"/api/login"} method="post">
                    <div className="form-group">
                        <label htmlFor="login-name">Your name</label>
                        <input type="text" className="form-control" id="login-name" name="name" placeholder="Your Name" />
                    </div>
                    <div className="form-group">
                        <label htmlFor="login-password">Password</label>
                        <input type="password" className="form-control" id="login-password" name="password" placeholder="Your Password" />
                    </div>
                    <div className="form-group">
                        <button type="submit" className="btn btn-primary">Login</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
}