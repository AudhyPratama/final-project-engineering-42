import React, { useEffect, useState, useRef } from "react";
import { Link, Router } from "react-router-dom";

import ForgotPassword from "./ForgotPassword";
import SignupForm from "./SignupForm";

import userIcon from "../../../asset/icons/user.png";
import emailIcon from "../../../asset/icons/email.png";
import lockIcon from "../../../asset/icons/padlock.png";
import readingIcon from "../../../asset/icons/reading.png";

import './LoginForm.css';

function LoginFormComponent () {

    const email = useRef();
    const password = useRef();
    const handleSubmit = () => {
        if (email.current.value == "dito@gmail.com" && password.current.value == "dito7654" ) {
            localStorage.setItem ("");
        }
    }

    return (
        <div className="main">
            <div className="sub-main">
                <div>
                <h1>Welcome back <span><img src={readingIcon} alt="reading icon" /></span></h1>
                    <div className="imgs">
                        <div className="container-image"> 
                            <img src={userIcon} alt="profile" className="profile"/>
                        </div>
                    </div>

                    <div>
                        <h2>Sign in page</h2>
                        <form onSubmit={handleSubmit}>
                            <div>
                                <h3>Username</h3>
                                <img src={emailIcon} alt="email" className="email"/>
                                <input type="text" placeholder="loremipsum@gmail.com" className="name" ref={email}/>
                            </div>
                            <div className="second-input">
                                <h3>Password</h3>
                                <img src={lockIcon} alt="pass" className="email"/>
                                <input type="password" placeholder="password" className="name"ref={password}/>
                            </div>
                            <div className="login-button">
                                <button type="submit" >Login</button>
                            </div>
                        </form>
                        <div className="link">
                            <a href="#">Forgot password ?</a> Or <a href="#">Sign Up</a>
                            {/* <Router exact path="/forgot-password" component={ForgotPassword}/>
                            <Router exact path="/sign-up" component={SignupForm}/> */}
                            
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default LoginFormComponent;