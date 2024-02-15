import React from "react";
import { useState, useEffect } from "react";
import { FaSearch } from "react-icons/fa";
import menuIcon from "../icons/menu.svg";
import backwardIcon from "../icons/backward.svg";
import { useNavigate } from "react-router-dom";
import MessageList from "./MessageList";
import MessageForm from "./MessageForm";

const ChatroonHeader = () => {
  const navigate = useNavigate();
  const handleClickToChatRoom = () => {
    navigate("/message");
  };
  return (
    <div>
      <div style={{ display: "flex", position: "relative" }}>
        <button
          style={{
            background: "white",
            border: "white",
            margin: "0px",
            position: "absolute",
            left: "5%",
            top: "20px",
          }}
          onClick={handleClickToChatRoom}
        >
          <img src={backwardIcon} alt="Backward Icon" />
        </button>
        <p
          style={{
            fontSize: "12px",
            position: "absolute",
            left: "15%",
            top: "15px",
            fontWeight: "bold",
          }}
        >
          대전광역시 유성구 전체
        </p>
        <button
          style={{
            background: "white",
            border: "white",
            margin: "0px",
            position: "absolute",
            right: "14%",
            top: "20px",
          }}
        >
          <FaSearch size={25} color="FF772A" />
        </button>
        <button
          style={{
            background: "white",
            border: "white",
            margin: "0px",
            position: "absolute",
            top: "17px",
            right: "5%",
          }}
        >
          <img src={menuIcon} alt="Menu Icon" />
        </button>
        <div
          style={{
            borderTop: "1px solid #A3A5A8",
            position: "absolute",
            top: "70px",
            width: "100%",
          }}
        ></div>
      </div>
    </div>
  );
};

export default ChatroonHeader;
