// MODULE
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useEffect } from 'react';

// IMAGE
import userIcon from '../../../asset/icons/user.png';
import readingIcon from '../../../asset/icons/reading.png';
import emailIcon from '../../../asset/icons/email.png';
import lockIcon from '../../../asset/icons/padlock.png';
import nameIcon from '../../../asset/icons/user-white.png';

// CSS
import './ProfileForm.css';

const ProfileForm = () => { 

    // update profile
    const [state, setState] = useState({
        name: "",
        email: "",
        password: "",
        role: ""
    });

    const handleChange = (prop) => (event) => {
        setState({ ...state, [prop]: event.target.value });
    }

    const navigate = useNavigate(); 

    // simpan profile di local storage
    const saveProfile = () => {
        localStorage.setItem('name', state.name);
        localStorage.setItem('email', state.email);
        localStorage.setItem('password', state.password);
        localStorage.setItem('role', state.role);
    };

    // tampilkan profile di local storage
    useEffect(() => {
        setState({
            name: localStorage.getItem('name'),
            email: localStorage.getItem('email'),
            password: localStorage.getItem('password'),
            role: localStorage.getItem('role')
        });
    }
    , []);


    return (
        <div className='main'>
            <div className='sub-main'>
                <div>
                    <h1 className='welcome'><span><img src={userIcon} alt="userIcon"></img></span>Hello</h1>
                    
                    <div>
                        <h2>Your Profile</h2>
                        <form >
                            <div>
                                
                                <h3 className='email-sign-in-text'></h3>
                                <img src={nameIcon} alt="name" className="email-sign-in-icon"/>
                                <input type="text" placeholder="Name" className="login-name" onChange={handleChange('name')} value={state.name} />
                            </div>
                            <div className="second-input">
                                <h3 className='password-sign-in-text'></h3>
                                <img src={emailIcon} alt="email" className="password-sign-in-icon"/>
                                <input type="text" placeholder="email" className="login-name" onChange={handleChange('email')} value={state.email} />
                            </div>
                            <div className="second-input">
                                <h3 className='password-sign-in-text'></h3>
                                <img src={lockIcon} alt="pass" className="password-sign-in-icon"/>
                                <input type="password" placeholder="password" className="login-name" onChange={handleChange('password')} value={state.password}/>
                            </div>
                            <div className="second-input">
                                <h3 className='password-sign-in-text'></h3>
                                <img src={nameIcon} alt="name" className="password-sign-in-icon"/>
                                <input type="text" placeholder="Role" className="login-name" onChange={handleChange('role')} value={state.role} />
                            </div>
                            <div className="login-button">
                                <button type="submit" className="login-btn" onClick = {saveProfile}>Save
                                </button>
                            </div>
                        </form>
                        <div className="login-button">
                            <button 
                                className="login-btn" 
                                href="/dashboard" 
                                onClick={() => navigate('/dashboard')
                            }>
                                Back To Dashboard
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div> 
    );
}

export default ProfileForm;
