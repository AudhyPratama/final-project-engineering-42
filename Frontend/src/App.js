// MODULES
import React from 'react';
import { Routes, Route } from "react-router-dom";

// FILES
import Login from './page/login/Login';
import Dashboard from './page/dashboard/Dashboard';
import SignUp from './page/signup/SignUp';
import ForgotPassword from './page/forgot-password/ForgotPassword';
import Profile from './page/profile/profile';
import Payment from './page/Payment/payment';
import Admin from './page/adminPage/admin';

// CSS
import './App.css';

function App() {
  return (
    <>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/sign-up" element={<SignUp />} />
        <Route path="/forgot-password" element={<ForgotPassword/>} />
        <Route path="/profile" element={<Profile/>} />
        <Route path="/payment" element={<Payment/>} />
        <Route path="/admin-page" element={<Admin/>} />
      </Routes>
    </>
  );
}

export default App;
