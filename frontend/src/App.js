import React, { useEffect } from "react";
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

function App() {
  const user = useSelector(selectUser);
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
  // console.log(user);

  if (!user) {
    return <Login />;
  } else {
    return (
      <div className="App">
        <QuoraHeader />
        <Routes>
          {/* <Route path="/" element={<Login />} /> */}
          <Route exact path="/" element={<Quora />} />
          <Route exact path="/profile" element={<Profile user={user} />} />
        </Routes>
        {/* <Quora /> */}
        {/* <Profile /> */}
      </div>
    );
  }
}

export default App;
