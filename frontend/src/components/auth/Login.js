import React from "react";
import "./Login.css";
import { signInWithPopup } from "firebase/auth";
import { auth, provider } from "../../firebase";
import axios from "axios";
import logo1 from "../images/kora1.png";

function Login(props) {
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
            .then((response) => {
              localStorage.setItem('Login', true);
              props.setIsLoggedIn(true);
              // console.log("response", response);
            })
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
            .then((response) => {
              localStorage.setItem('Login', true);
              props.setIsLoggedIn(true);
              props.setNewuser(true);
              localStorage.setItem('NewUser', true);
            })
            .catch((error) => {
              console.log(error);
            });
        }
      })
  };

  return (
    <div className="login-container">
      <div className="login-content">
      <img
          src={logo1}
          alt="logo"
          style={{
            borderRadius: "20px",
            width: "500px",
            marginBottom: "-20px",
          }}
        />
        <button onClick={handleSubmit} className="btn-login">
          Login to continue
        </button>
      </div>
    </div>
  );
}

export default Login;
