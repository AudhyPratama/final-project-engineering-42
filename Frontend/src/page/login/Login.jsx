import React from 'react';

import LoginForm from './components/LoginForm';

import './Login.css';

const Login = () => {
    return (
        <div className='login-page'>
            <h1 className='App-title-1'>Pes<span>Buk</span></h1>
            <LoginForm/>
        </div>
    );
}

export default Login;
