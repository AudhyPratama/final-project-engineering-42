import React, { useEffect, useState } from "react";
// import { Link } from "react-router-dom";

import userIcon from "../asset/icons/user.png";
import emailIcon from "../asset/icons/email.png";
import lockIcon from "../asset/icons/padlock.png";
import readingIcon from "../asset/icons/reading.png";

import '../asset/css/LoginForm.css';

function LoginFormComponent () {
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
                        <div>
                            <img src={emailIcon} alt="email" className="email"/>
                            <input type="text" placeholder="loremipsum@gmail.com" className="name"/>
                        </div>
                        <div className="second-input">
                            <img src={lockIcon} alt="pass" className="email"/>
                            <input type="password" placeholder="password" className="name"/>
                        </div>
                        <div className="login-button">
                            <button type="submit" ><a href="#"></a></button>
                        </div>      
                        <div className="link">
                            <a href="#">Forgot password ?</a> Or <a href="#">Sign Up</a>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default LoginFormComponent;