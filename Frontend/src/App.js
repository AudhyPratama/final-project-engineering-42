import React from 'react';
// import { Navigate } from "react-router-dom";
import "./App.css";
import DashboardUi from './page/dashboard/Dashboard';
import LoginUi from "./page/login/Login";
import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";



function App() {
  return (
    <Routes>
      <Route path="/" element={<LoginUi />} />
      <Route path="/dashboard" element={<DashboardUi />} />
    </Routes>
    // <div className='App'>
    //   <LoginUi />
    //   <Navigate to="/dashboard"/>
    //   {/* <DashboardUi /> */}
    // </div> 
  );
}

export default App;
