import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import "./App.css";
import Login from "./components/auth/Login";
import Quora from "./components/Quora";
import { login, selectUser } from "./feature/userSlice";
import { auth } from "./firebase";
import { onAuthStateChanged } from "firebase/auth";
import Profile from "./components/Profile";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import QuoraHeader from "./components/QuoraHeader";
import NewCategories from "./components/NewCategories";
import Answer from "./components/Answer";

function App() {
  const user = useSelector(selectUser);
  const [showFade, setShowFade] = useState(false);
  const [newUser, setNewuser] = useState(false);
  const dispatch = useDispatch();

  useEffect(() => {
    onAuthStateChanged(auth, (authUser) => {
      if (authUser) {
        dispatch(
          login({
            userName: authUser.displayName,
            photo: authUser.photoURL,
            email: authUser.email,
            uid: authUser.uid,
          })
        );
        console.log("AuthUser", authUser);
      }
    });
  }, [dispatch]);

  if (!user) {
    return <Login setNewuser={setNewuser}/>;
  } else
  if (newUser) {
    return <NewCategories setNewuser={setNewuser}/>
  } else {
    return (
      <div className="App">
        <QuoraHeader setShowFade={setShowFade} />
        <Routes>
          <Route
            exact
            path="/"
            element={<Quora showFade={showFade} setShowFade={setShowFade} />}
          />
          <Route exact path="/profile" element={<Profile user={user} />} />

          <Route exact path="/answers" element={<Answer />} />
        </Routes>
      </div>
    );
  }
}

export default App;
