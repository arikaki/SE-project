import React from "react";
import "./Login.css";
import { signInWithPopup } from "firebase/auth";
import { auth, provider } from "../../firebase";
import axios from "axios";

function Login() {
  const handleSubmit = async () => {
    await signInWithPopup(auth, provider)
      .then((result) => {
        const { user } = result;
        const { createdAt, lastLoginAt } = user.metadata;
        if (Math.abs(createdAt - lastLoginAt) > 10) {
          axios.post('http://localhost:8000/login', {
            UserName: user.displayName,
            Password: 'password'
          }, {
            "withCredentials": true,
            "Access-control-Allow-Origin": "http://localhost:8000"
          })
            .then((response) => console.log("response", response))
            .catch((error) => {
              console.log(error);
            });
        }
        else {
          axios.post('http://localhost:8000/api/user/insert', {
            Name: user.displayName,
            Email: user.email,
            Password: "password",
            Username: user.displayName
          }, {
            "withCredentials": true,
            "Access-control-Allow-Origin": "http://localhost:8000"
          })
            .then((response) => console.log("response", response))
            .catch((error) => {
              console.log(error);
            });
        }
      })
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