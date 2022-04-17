import React from "react";
import "./Login.css";
import { signInWithPopup } from "firebase/auth";
import { auth, provider } from "../../firebase";
import axios from "axios";

function Login() {
  const handleSubmit = async () => {
    await signInWithPopup(auth, provider)
      .then((result) => {
        console.log(result.user.displayName);
        axios.post('http://localhost:8000/login'), {
         UserName: user.displayName,
         Password: 'kirthi',
        },{
          "Access-control-Allow-Origin": "*"
        }

      })
      .catch((error) => {
        console.log(error);
      });
  };
  return (
    <div className="login-container">
      <div className="login-content">
        {/* <img
          src="https://video-public.canva.com/VAD8lt3jPyI/v/ec7205f25c.gif"
          alt="logo"
        /> */}
        <button onClick={handleSubmit} className="btn-login">
          Login to continue
        </button>
      </div>
    </div>
  );
}

export default Login;