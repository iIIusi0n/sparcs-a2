import React from "react";
import homeIcon from "../icons/Frame 14(1).svg";
import pinIcon from "../icons/map (2).svg";
import messageIcon from "../icons/chat.svg";
import profileIcon from "../icons/c.png";
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
        bottom: "20%",
        width: "100%",
        marign: "0",
      }}
    >
      <div
        style={{
          backgroundColor: "white",
          position: "absolute",
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          padding: "15px 10px 10px",
          paddingBottom: "20px",
          width: "100%",
          height: "5vh",
        }}
      >
        <div
          style={{
            position: "absolute",
            bottom: "20px",
            left: "80px",
            display: "flex",
            flexDirection: "column",
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
              fontSize: "14px",
              fontFamily: "Pretendard",
              position: "relative",
              margin: "0",
              left: "14px",
            }}
          >
            홈
          </p>
        </div>
        <div
          style={{
            position: "absolute",
            bottom: "20px",
            left: "160px",
            display: "flex",
            flexDirection: "column",
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
              fontSize: "14px",
              fontFamily: "Pretendard",
              position: "relative",
              margin: "0",
              left: "10px",
            }}
          >
            PIN
          </p>
        </div>
        <div
          style={{
            position: "absolute",
            bottom: "20px",
            left: "240px",
            display: "flex",
            flexDirection: "column",
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
              fontSize: "14px",
              fontFamily: "Pretendard",
              position: "relative",
              left: "2px",
              margin: "0",
              color: "#FF772A",
            }}
          >
            채팅
          </p>
        </div>

        <div
          style={{
            position: "absolute",
            bottom: "22px",
            left: "320px",
            display: "flex",
            flexDirection: "column",
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
              fontSize: "14px",
              fontFamily: "Pretendard",
              position: "relative",
              margin: "0",
              top: "4px",
              left: "2px",
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
