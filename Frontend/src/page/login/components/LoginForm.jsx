// MODULE
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';

// IMAGE
import userIcon from '../../../asset/icons/user.png';
import readingIcon from '../../../asset/icons/reading.png';
import emailIcon from '../../../asset/icons/email.png';
import lockIcon from '../../../asset/icons/padlock.png';

// CSS
import './LoginForm.css';

const LoginForm = () => { 
    const [state, setState] = useState({
        email: '',
        password: ''
    });

    const handleChange = (prop) => (event) => {
        setState({ ...state, [prop]: event.target.state });
    };

    const navigate = useNavigate();

    const handleSubmit = async (event) => {
        event.preventDefault();
        // console.log(values);
        await axios
            .post(
                'http://localhost:8080/api/user/login',
                {
                    email: state.email,
                    password: state.password,
                },
                {
                    Headers: {
                        'Content-Type': 'application/json',
                    },
                }    
            )
            .then((response) => {
                console.log("axios", response);
                localStorage.setItem("token", response.data.token);
                console.log(response.data.token);
                window.location.href = "/dashboard";
            })
            .catch (function (error) {
                console.log(error);
            })
    }

    return (
        <div className='main'>
            <div className='sub-main'>
                <div>
                    <h1 className='welcome'>Welcome Back <span><img src={readingIcon} alt="reading icon"></img></span></h1>
                    <div className='imgs'>
                        <div className='container-image'>
                            <img src={userIcon} alt="profile" className="profile"/>
                        </div>
                    </div>

                    <div>
                        <h2>Sign in page</h2>
                        <form onSubmit={handleSubmit}>
                            <div>
                                <h3 className='email-sign-in-text'>Email</h3>
                                <img src={emailIcon} alt="email" className="email-sign-in-icon"/>
                                <input type="text" placeholder="loremipsum@gmail.com" className="login-name" onChange={handleChange('email')}/>
                            </div>
                            <div className="second-input">
                                <h3 className='password-sign-in-text'>Password</h3>
                                <img src={lockIcon} alt="pass" className="password-sign-in-icon"/>
                                <input type="password" placeholder="password" className="login-name" onChange={handleChange('password')}/>
                            </div>
                            <div className="login-button">
                                <button className="login-btn" type="submit" >Login</button>
                            </div>
                        </form>
                        <div className="link">
                            <a 
                                className='link-forgot-pw' 
                                href="/forgot-password" 
                                onClick={() => 
                                    navigate('/forgot-password')
                            }>
                                Forgot password ?
                            </a> Or 
                            <a 
                                className='link-sign-up' 
                                href="/sign-up" 
                                onClick={() => 
                                    navigate('/sign-up')
                                }
                            >
                                Sign Up
                            </a>
                        </div>
                    </div>
                </div>
            </div>
        </div> 
    );
}

export default LoginForm;
