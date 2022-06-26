import React from 'react';

import ProfileForm from './Components/ProfileForm';

import './profile.css';

const Profile = () => {
    return (
        <div className='profile-page'>
            <h1 className='App-title-1'>Pes<span>Buk</span></h1>
            <ProfileForm/>
        </div>
    );
}

export default Profile;
