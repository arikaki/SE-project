import React, { useEffect, useState } from "react";
import "./App.css";
import Header from "./Components/Header";
import QuestionBox from "./Components/QuestionBox";
import Question from "./Question";
import axios from "axios";
import QuestionList from "./Components/QuestionList";
const App = () => {
  const [showAskQuestion, setShowAskQuestion] = useState(false);
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
    <div>
      {/* <Router className="App"> */}
        {/* <Switch> */}
          {/* <Route path="/user-signup"> */}
            <SignUp />
          {/* </Route> */}
          {/* <Route path="/user-signinside"> */}
            <SignInSide />
          {/* </Route> */}
          {/* <Route path="/"> */}
            {/* <QuestionBox auth_status={auth_status} profile={profile} />
          <QuestionList /> */}
          {/* </Ro/ute> */}
        {/* </Switch> */}
      {/* </Router> */}
      <div className="App">

        <Header onAsk={onAsk} />
        {showAskQuestion ? <div style={{ marginTop: "10%" }}>
          <QuestionBox closeQuestion={closeQuestion} />
          <QuestionList />
        </div> : <Question />}
      </div>
    </div>
  );
};

export default App;