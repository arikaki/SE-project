import React from "react";
import "./Login.css";
import { signInWithPopup } from "firebase/auth";
import { auth, provider } from "../../firebase";
import axios from "axios";
import logo1 from "../images/kora1.png";

function Login() {
  const handleSubmit = async () => {
    await signInWithPopup(auth, provider)
      .then((result) => {
        console.log(result.user.displayName);
      })
      .catch((error) => {
        console.log(error);
      });
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
