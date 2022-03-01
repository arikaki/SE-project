import React from "react";
import HomeIcon from "@material-ui/icons/Home";
import NotificationsIcon from "@material-ui/icons/Notifications";
import CustomButton from "./CustomButton";
import ReorderIcon from '@mui/icons-material/Reorder';
import Avatar from "@material-ui/core/Avatar";
import Search from './Search';
import Logo from "../images/logo.jpeg";
import BasicMenu from "./BasicMenu";

import "../StyleSheet/Header.css";
const Header = ({ onAsk }) => {
  const location = window.location.href;

  return (
    <div className="Header">
      <div className="Header__left">
        
        <div
          className={`left__Home ${
            location === "http://localhost:3000/" ? "current-location" : null
          }`}
        >
          <CustomButton onClick={()=>{window.alert("Home")}}>
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
          <CustomButton onClick={()=>{window.alert("Notification")}}>
            <NotificationsIcon />
            <h4>Notifications</h4>
          </CustomButton>
          <CustomButton onClick={()=>{window.alert("Categories")}}>
            <ReorderIcon />
            <h4>Categories</h4>
          </CustomButton>
          
        </div>
      </div>
      <div className="left__logo">
        
        </div>
      <div className="Header__center">
        <Search />
      </div>
      <div className="Header__right">
        <div className="right__user">
          <BasicMenu>
            <Avatar alt="User Profile" />
          </BasicMenu>
        </div>
        <button className="right-btn" onClick={onAsk}>Ask Anything</button>
      </div>
    </div>
  );
};

export default Header;