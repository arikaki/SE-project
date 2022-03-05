import React, { useEffect, useState } from "react";
import "./App.css";
//import "./StyleSheet/profile.css"
import Header from "./Components/Header";
import QuestionBox from "./Components/QuestionBox";
import Question from "./Question";
import axios from "axios";
import QuestionList from "./Components/QuestionList";
import SignUp from "./Components/SignUp";
import SignInSide from "./Components/SignInSide";
import Profile from "./Components/Profile";

const App = () => {
  const [showAskQuestion, setShowAskQuestion] = useState(false);
  const [showSignup, setShowSignup] = useState(false);
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const onAsk = () => {
    setShowAskQuestion(true);
  }
  const closeQuestion = () => {
    setShowAskQuestion(false);
  }

  // useEffect(() => {
  //   const url = "http://localhost:3000/api/isUserLoggedIn";
  //   axios
  //     .get(url, { withCredentials: true })
  //     .then((response) => {
  //       console.log(response);
  //       setAuthStatus(response.data.auth_status);
  //       setImage(response.data.profileImage);
  //     })
  //     .catch((error) => {
  //       console.log(error);
  //     });
  // }, []);

  return (
      <div className="App">
        {/* {showSignup ? <SignUp setShowSignup={setShowSignup} /> : !isLoggedIn ? <SignInSide setShowSignup={setShowSignup} /> :
          <>
            <Header onAsk={onAsk} />

            {showAskQuestion ? <div style={{ marginTop: "10%" }}>
              <QuestionBox closeQuestion={closeQuestion} />
              <QuestionList />
            </div> : <Question />}
          </>} */}
          <Profile/>
          
         
      </div>
  );
};

export default App;