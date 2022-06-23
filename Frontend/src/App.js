import React from 'react';
// import './App.css';
import { Routes, Route } from "react-router-dom";
import Login from './page/login/Login';
import Dashboard from './page/dashboard/Dashboard';
import SignUp from './page/signup/SignUp';

function App() {
  return (
    <>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/sign-up" element={<SignUp />} />
        {/* <Route path="/forgot-password" element={<ForgotPassword/>} /> */}
      </Routes>
    </>
  );
}

export default App;
