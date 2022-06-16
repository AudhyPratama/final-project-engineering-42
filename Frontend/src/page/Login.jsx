import LoginFormComponent from "../components/LoginForm";

import "../asset/css/Login.css";

function LoginUi () {
    return (
        <div className="container">
            <h1 className="title-1">Pes<span>Buk</span></h1>
            <LoginFormComponent />
        </div>
    );
    
}

export default LoginUi;