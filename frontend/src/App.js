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
import CategoryPage from "./components/CategoryPage";

function App() {
  const user = useSelector(selectUser);
  const [showFade, setShowFade] = useState(false);
  const [newUser, setNewuser] = useState(localStorage.getItem('NewUser') == "true");
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const dispatch = useDispatch();
  
  useEffect(() => {
    setIsLoggedIn(localStorage.getItem('Login') == 'true');
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
      }
    });
  }, [dispatch]);


  if (!isLoggedIn) {
    return <Login setIsLoggedIn={setIsLoggedIn} setNewuser={setNewuser}/>;
  } else
  if (newUser) {
    return <NewCategories setNewuser={setNewuser}/>
  } else {
    return (
      <div className="App">
        <QuoraHeader setIsLoggedIn={setIsLoggedIn} setShowFade={setShowFade} />
        <Routes>
          <Route exact path="/" element={<Quora showFade={showFade} setShowFade={setShowFade} />}
          />
          <Route exact path="/profile" element={<Profile user={user} showFade={showFade} setShowFade={setShowFade} setIsLoggedIn={setIsLoggedIn}/>} />
          <Route exact path="/categories" element={<CategoryPage user={user} notRegister={true} showFade={showFade} setShowFade={setShowFade}/>}/>
        </Routes>
      </div>
    );
  }
}

export default App;
