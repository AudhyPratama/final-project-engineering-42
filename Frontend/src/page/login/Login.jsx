import LoginFormComponent from "../login/components/LoginForm";
import ForgotPassword from "./components/ForgotPassword";
import SignupForm from "./components/SignupForm";

import "./Login.css";

function LoginUi () {
    return (
        <div className="container">
            <h1 className="title-1">Pes<span>Buk</span></h1>
            <LoginFormComponent />
        </div>
    );
    
}

export default LoginUi;