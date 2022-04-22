import React from "react";
import Feed from "./Feed";
import Sidebar from "./Sidebar";
import Widget from "./Widget";
import "./css/Quora.css";

function Quora(props) {
  return (
    <div className={`quora${props.showFade? " fade-search": ""}`} onClick={() => props.setShowFade(false)}>
      <div className="quora__contents">
        <div className="quora__content">
          <Sidebar />
          <Feed />
          <Widget />
        </div>
      </div>
    </div>
  );
}

export default Quora;
