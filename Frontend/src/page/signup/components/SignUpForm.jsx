// MODULE
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';

// IMAGE
import userIcon from '../../../asset/icons/user-white.png';
import emailIcon from '../../../asset/icons/email.png';
import lockIcon from '../../../asset/icons/padlock.png';
import backIcon from '../../../asset/icons/left-arrow.png';

// CSS
import './SignUpForm.css';

const SignUpForm = () => {

    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [role, setrole] = useState('');

    const navigate = useNavigate();

    const handleSubmit = (e) => {
        e.preventDefault();
        axios.post('http://localhost:8080/api/registrasi', {
            name,
            email,
            password,
            role
        }).then(res => {
            console.log(res);
            navigate('/');
        }
        ).catch(err => {
            console.log(err);
        });
    }

    return (
        <div className='sign-up-component'>
            <div className='sign-up-container'>
                <div>
                    <div className='back'>
                        <a 
                            href="/" 
                            className='back-link'
                            onClick={() => 
                                navigate('/')
                            }
                        >
                            <img 
                                src={backIcon} 
                                alt='back icon'
                                className='back-icon' 
                            />    
                            <span className='back-button-text'>
                                Back
                            </span>    
                            
                        </a>
                    </div>
                    <div className='sign-up-title'>
                        <h1 className='App-title-2'>Sign up to <span className='pes'>Pes</span><span className='buk'>Buk</span></h1>
                    </div>

                    <div className='sign-up-form-container'>
                        <form className='sign-up-form' onSubmit={handleSubmit}>
                            <div>
                                <h3 className='username-sign-up-text'>Name</h3>
                                <img src={userIcon} alt="user icon" className='sign-up-user-icon' />
                                <input type="text" placeholder="Name" value={name} className="signup-name" onChange={(e) => setName(e.target.value)} />
                            </div>
                            <div>
                                <h3 className='email-sign-up-text'>Email</h3>
                                <img src={emailIcon} alt="email icon" className='sign-up-email-icon' />
                                <input type="text" placeholder="Email" value={email} className="signup-name" onChange={(e) => setEmail(e.target.value)} />
                            </div>
                            <div>
                                <h3 className='password-sign-up-text'>Password</h3>
                                <img src={lockIcon} alt="password icon" className='sign-up-password-icon' />
                                <input type="password" placeholder="Password" value={password} className="signup-name" onChange={(e) => setPassword(e.target.value)} />
                            </div>
                            <div>
                                <h3 className='username-sign-up-text'>Role</h3>
                                <img src={userIcon} alt="user icon" className='sign-up-user-icon' />
                                <input type="text" placeholder="Role" value={role} className="signup-name" onChange={(e) => setrole(e.target.value)} />
                            </div>
                            <div className="sign-up-button">
                                <button className="sign-in-btn" type="submit" >Signup</button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default SignUpForm;