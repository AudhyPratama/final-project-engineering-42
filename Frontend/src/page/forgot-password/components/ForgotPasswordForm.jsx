// MODULE
import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

// IMG and icons
import emailIcon from '../../../asset/icons/email.png';
import lockIcon from '../../../asset/icons/padlock.png';
import cautionMark from '../../../asset/icons/caution.png';

// CSS
import './ForgotPasswordForm.css';

const ForgotPasswordForm = () => {

    const [state, setState] = useState({
        email: '',
        password: '',

    });

    const navigate = useNavigate();

    const handleChange = (e) => {
        setState({...state, [e.target.name]: e.target.value });
    }

    const handleSubmit = (e) => {
        e.preventDefault();
        axios.post('http://localhost:8080/api/login/forgot-password', state)
            .then(res => {
                console.log(res);
                console.log(res.data);
                navigate('/');
            }
            )
            .catch(err => {
                console.log(err);
            });
    };

    return (
        <div className='forgot-password-component'>
            <div className='forgot-password-container'>
                <div>
                    <h1 className='forgot-password-text'>Forgot password 
                        <span className='question-mk'>?</span>
                    </h1>
                    <div className='forgot-password-images'>
                        <div className='container-ex-image'>
                            <img 
                                src={cautionMark} 
                                alt="exclamation mark" 
                                className="exclamation-mk"
                            />
                        </div>
                    </div>

                    <div>
                        <div className='forgot-password-description'>
                            <h2 className='new-email-password-description'>
                                <p className='description'>Reset your account password by entering your registered email and new password</p>
                            </h2>
                        </div>
                        <div className='forgot-password-form-container'>
                            <form  className='forgot-password-form' onSubmit={handleSubmit}>
                                <div>
                                    <h3 className='email-forgot-password-text'>Email</h3>
                                    <img src={emailIcon} alt="email icon" className="email-forgot-password-icon"/>
                                    <input type="text" name="email" placeholder="Enter your email" onChange={handleChange} />
                                </div>
                                <div>
                                    <h3 className='password-forgot-password-text'>New Password</h3>
                                    <img src={lockIcon} alt="password icon" className="password-forgot-password-icon"/>
                                    <input type="password" name="password" placeholder="Enter your new password" onChange={handleChange} />
                                </div>
                                <div className="submit-button">
                                    <button className="submit-btn" type="submit" >Submit</button>
                                </div>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        </div> 
    );
}

export default ForgotPasswordForm;

