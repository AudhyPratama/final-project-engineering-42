// MODULE
import React from 'react';

// IMAGE
import userIcon from '../../../asset/icons/user-white.png';
import emailIcon from '../../../asset/icons/email.png';
import lockIcon from '../../../asset/icons/padlock.png';

// CSS
import './SignUpForm.css';

const SignUpForm = () => {
    const post = 1;
    // const handleChange ='' ;
    return (
        <div className='sign-up-component'>
            <div className='sign-up-container'>
                <div>
                    <div className='back'>
                        <a href='#'>Back</a>
                    </div>
                    <div className='sign-up-title'>
                        <h1> Sign up page </h1>
                    </div>

                    <div className='sign-up-form-container'>
                        <form className='sign-up-form' onSubmit={post}>
                            <div>
                                <h3>Username</h3>
                                <img src={userIcon} alt="user icon" className='user-icon' />
                                <input type="text" placeholder="Username" className="signup-name" /*onChange={handleChange('username')}*//>
                            </div>
                            <div>
                                <h3>Email</h3>
                                <img src={emailIcon} alt="email icon" className='email-icon' />
                                <input type="text" placeholder="Email" className="signup-name" /*onChange={handleChange('email')}*//>
                            </div>
                            <div>
                                <h3>Password</h3>
                                <img src={lockIcon} alt="password icon" className='password-icon' />
                                <input type="password" placeholder="Password" className="signup-name" /*onChange={handleChange('password')}*//>
                            </div>
                            <div className="sign-in-button">
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
