import React, { useState } from "react";
import Header from "../components/header/header.jsx";
import NavigationBar from "../components/navigation/NavigationBar.jsx";
import Title from "../components/messageComponents/messageTitle.jsx";
import finIcon from "../components/icons/fin.svg";
import MessageMove from "../components/messageComponents/MessageMove.jsx";
import NavigationBarChat from "../components/navigation/NavigationBarChat.jsx";

const MessageMain = (props) => {
  const id = props;
  const [latitude, setLatitude] = useState(null);
  const [longitude, setLongitude] = useState(null);
  const [error, setError] = useState(null);

  const getLocation = () => {
    if ("geolocation" in navigator) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          setLatitude(position.coords.latitude);
          setLongitude(position.coords.longitude);
          setError(null);
        },
        (error) => {
          setError(error.message);
        }
      );
    } else {
      setError("Geolocation is not supported in this browser.");
    }
  };
  return (
    <div>
      <div
        style={{
          position: "relative",
          display: "flex",
          flexDirection: "column",
          paddingLeft: "5%",
          paddingRight: "5%",
        }}
      >
        <Header></Header>
        <Title></Title>
        <div
          style={{
            display: "flex",
            position: "absolute",
            padding: "0px",
            top: "600%",
            left: "8%",
          }}
        >
          <img src={finIcon} alt="Fin Icon" />
          <p style={{ fontWeight: "bold", fontSize: "20px" }}>
            참여 가능한 채팅방
          </p>
          <p
            style={{
              fontWeight: "bold",
              fontSize: "8px",
              marginLeft: "20px",
              position: "relative",
              top: "25px",
              left: "15%",
              color: "#A3A5A8",
            }}
          >
            위치 기반으로 추천해요
          </p>
        </div>
        <div style={{ position: "absolute", top: "880%", left: "10%" }}>
          <MessageMove props={id}></MessageMove>
        </div>
      </div>
      <NavigationBarChat></NavigationBarChat>
      <NavigationBarChat></NavigationBarChat>
      <NavigationBarChat></NavigationBarChat>
    </div>
  );
};
export default MessageMain;
