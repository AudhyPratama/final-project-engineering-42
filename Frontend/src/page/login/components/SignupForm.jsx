import React, { useState } from "react";

import LoginFormComponent from "./LoginForm";

import './SignupForm.css';

function SignupForm () {

    const [emailRegist, setEmailRegist] = useState('');
    const [passRegist, setPassRegist] = useState('');

    return (
        <div className='sign-up-form'>
            <h1>ini halaman Signup form</h1>
{/*             
            <input 
                type="text" onChange={(e) => {
                    setEmailRegist(e.target.value);
                }}
            />
            <input 
                type="text" onChange={(e) => {
                    setPassRegist(e.target.value);
                }}
            />
*/}
        </div>
    );
}

export default SignupForm;