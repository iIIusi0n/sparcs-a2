import React from "react";
import homeIcon from "../icons/Frame 14(1).svg";
import pinIcon from "../icons/map (2).svg";
import messageIcon from "../icons/chat.svg";
import profileIcon from "../icons/Frame 15.svg";
import { useNavigate } from "react-router-dom";

const NavigationBarChat = () => {
  const navigate = useNavigate();
  const handleClickToMap = () => {
    // 버튼을 클릭하면 '/about' 경로로 이동
    navigate("/map");
  };
  const handleClickToHome = () => {
    // 버튼을 클릭하면 '/about' 경로로 이동
    navigate("/");
  };
  const handleClickToChat = () => {
    // 버튼을 클릭하면 '/about' 경로로 이동
    navigate("/message");
  };
  return (
    <div
      style={{
        position: "fixed",
        bottom: " 11%",
        width: "100%",
        marign: "0",
      }}
    >
      <div
        style={{
          backgroundColor: "white",
          position: "relative",
          display: "flex",
          padding: "2%",
          width: "100%",
          height: "5vh",
        }}
      >
        <div
          style={{
            display: "flex",
            flexDirection: "column",
            position: "relative",
            left: "15%",
          }}
        >
          <button
            style={{
              background: "white",
              border: "white",
              margin: "0px",
            }}
            onClick={handleClickToHome}
          >
            <img src={homeIcon} alt="Home Icon" />
          </button>
          <p
            style={{
              position: "absolute",
              fontSize: "14px",
              margin: "0",
              bottom: "0%",
              left: "35%",
              fontFamily: "Pretendard-Bold",
            }}
          >
            홈
          </p>
        </div>
        <div
          style={{
            display: "flex",
            flexDirection: "column",
            position: "relative",
            left: "25%",
          }}
        >
          <button
            style={{
              background: "white",
              border: "white",
              margin: "0px",
            }}
            onClick={handleClickToMap}
          >
            <img src={pinIcon} alt="Pin Icon" />
          </button>
          <p
            style={{
              position: "absolute",
              fontSize: "14px",
              margin: "0",
              bottom: "0%",
              left: "25%",

              fontFamily: "Pretendard-Bold",
            }}
          >
            PIN
          </p>
        </div>
        <div
          style={{
            display: "flex",
            flexDirection: "column",
            position: "relative",
            left: "36%",
          }}
        >
          <button
            style={{
              background: "white",
              border: "white",
              margin: "0px",
            }}
            onClick={handleClickToChat}
          >
            <img src={messageIcon} alt="Message Icon" />
          </button>
          <p
            style={{
              position: "absolute",
              fontSize: "14px",
              margin: "0",
              bottom: "0%",
              left: "16%",
              color: "#FF772A",

              fontFamily: "Pretendard-Bold",
            }}
          >
            채팅
          </p>
        </div>

        <div
          style={{
            display: "flex",
            flexDirection: "column",
            position: "relative",
            left: "45%",
          }}
        >
          <button
            style={{
              background: "white",
              border: "white",
              margin: "0px",
            }}
          >
            <img src={profileIcon} alt="Profile Icon" />
          </button>
          <p
            style={{
              position: "absolute",
              fontSize: "14px",
              margin: "0",
              bottom: "0%",
              left: "20%",

              fontFamily: "Pretendard-Bold",
            }}
          >
            마이
          </p>
        </div>
      </div>
    </div>
  );
};

export default NavigationBarChat;
