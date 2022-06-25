// MODULE
import React from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

// IMG and icons
import userIcon from '../../../asset/icons/user.png';
import emailIcon from '../../../asset/icons/email.png';
import lockIcon from '../../../asset/icons/padlock.png';
import cautionMark from '../../../asset/icons/caution.png';
// import backIcon from '../../../asset/icons/left-arrow.png';

// CSS
import './ForgotPasswordForm.css';

const ForgotPasswordForm = () => {
    const handleSubmit = 1;

    const navigate = useNavigate();

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
                                <p className='description'>Reset your account setting by entering your new email and password</p>
                            </h2>
                        </div>
                        <div className='forgot-password-form-container'>
                            <form  className='forgot-password-form' onSubmit={handleSubmit}>
                                <div>
                                    <h3 className='email-forgot-password-text'>New Email</h3>
                                    <img src={emailIcon} alt="email icon" className="email-forgot-password-icon"/>
                                    <input 
                                        type="text" 
                                        placeholder="loremipsum@gmail.com" 
                                        className="forgot-password-name" 
                                        // onChange={handleChange('email')}
                                    />
                                </div>
                                <div>
                                    <h3 className='password-forgot-password-text'>New Password</h3>
                                    <img src={lockIcon} alt="password icon" className="password-forgot-password-icon"/>
                                    <input 
                                        type="password" 
                                        placeholder="password" 
                                        className="forgot-password-name" 
                                        // onChange={handleChange('password')}
                                    />
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

