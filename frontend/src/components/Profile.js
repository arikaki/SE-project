import React from "react";
import "./css/Profile.css";
import { Tab, Tabs } from "react-bootstrap";
import { auth } from "../firebase";
import { signOut } from "firebase/auth";
import axios from "axios";
import { logout } from "../feature/userSlice";
import { useDispatch } from "react-redux";

export default function Profile({ user, showFade, setShowFade, setIsLoggedIn }) {
const dispatch = useDispatch();
  const deleteUser = () => {
    axios.get("http://localhost:8000/api/user/deleteuser", {
      "withCredentials": true,
      "Access-control-Allow-Origin": "http://localhost:8000"
    })
      .then((res) => {
        signOut(auth)
          .then(() => {
            localStorage.setItem('Login', false);
            setIsLoggedIn(false);
            dispatch(logout());
          })
          .catch(() => {
            console.log("error in logout");
          });
      })
      .catch((e) => {
        console.log(e);
      });
  }
  return (
    <div className={`container ${showFade ? " fade-search" : ""}`} onClick={() => setShowFade(false)}>
      <div className="main-body">
        <div className="row gutters-sm">
          <div className="col-md-4 mb-3">
            <div className="card">
              <div className="card-body">
                <div className="d-flex flex-column align-items-center text-center">
                  <img
                    src={user?.photo}
                    // src="https://bootdey.com/img/Content/avatar/avatar7.png"
                    alt="Image of User"
                    className="rounded-circle"
                    width={150}
                  />
                  <div className="mt-3 ">
                    <h4>{user?.userName}</h4>
                    <button className="btn btn-primary">Follow</button>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div className="col-md-8">
            <div className="card mb-3">
              <div className="card-body">
                <div className="row">
                  <div className="col-sm-3">
                    <h6 className="mb-0">Full Name</h6>
                  </div>
                  <div className="col-sm-9 text-secondary">
                    {user?.userName}
                  </div>
                </div>
                <hr />
                <div className="row">
                  <div className="col-sm-3">
                    <h6 className="mb-0">Email</h6>
                  </div>
                  <div className="col-sm-9 text-secondary">{user?.email}</div>
                </div>
                <hr />
                <div className="row">
                  <div className="col-sm-12">
                    <button className="btn btn-info" type='button' onClick={deleteUser}>
                      Delete
                    </button>
                  </div>
                </div>
              </div>
            </div>
            <div className="row gutters-sm">
              <Tabs
                defaultActiveKey="profile"
                id="uncontrolled-tab-example"
                className="mb-3"
              >
                <Tab eventKey="home" title="Answer">
                  {/* <Sonnet /> */}
                  <div>
                    <h1>Answer section</h1>
                    <p>All the answers questions, or count of it.</p>
                  </div>
                </Tab>
                <Tab eventKey="profile" title="Question">
                  {/* <Sonnet /> */}
                  <div>
                    <h1>Question Section </h1>
                    <p>All the questions asked by the user.</p>
                  </div>
                </Tab>
                <Tab eventKey="contact" title="Post">
                  {/* <Sonnet /> */}
                  <div>
                    <h1>Post Section</h1>
                    <p>All the posts done by the user</p>
                  </div>
                </Tab>
              </Tabs>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
