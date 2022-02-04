import React from "react";
import HomeIcon from "@material-ui/icons/Home";
import NotificationsIcon from "@material-ui/icons/Notifications";
import CustomButton from "./CustomButton";
import Avatar from "@material-ui/core/Avatar";
import Search from './Search';
import Logo from "../images/logo.jpeg";
import BasicMenu from "./BasicMenu";

import "../StyleSheet/Header.css";
const Header = ({ profile }) => {
  const location = window.location.href;

  return (
    <div className="Header">
      <div className="Header__left">
        
        <div
          className={`left__Home ${
            location === "http://localhost:3000/" ? "current-location" : null
          }`}
        >
          <CustomButton>
            <HomeIcon />
            <h4>Home</h4>
          </CustomButton>
        </div>
        <div
          className={`left__notifications ${
            location === "http://localhost:3000/notifications"
              ? "current-location"
              : null
          }`}
        >
          <CustomButton>
            <NotificationsIcon />
            <h4>Notifications</h4>
          </CustomButton>
        </div>
      </div>
      <div className="left__logo">
          <img src={Logo} alt="Company Logo" className="left__logoImage" />
        </div>
      <div className="Header__center">
        <Search />
      </div>
      <div className="Header__right">
        <div className="right__user">
          <BasicMenu>
            <Avatar src={profile} alt="User Profile" />
          </BasicMenu>
        </div>
        <button className="right-btn">Ask Anything</button>
      </div>
    </div>
  );
};

export default Header;