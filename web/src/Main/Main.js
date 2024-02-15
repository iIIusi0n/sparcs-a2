import React, { useState } from "react";
import styled from "styled-components";
import Header from "../components/header/header.jsx";
import NavigationBar from "../components/navigation/NavigationBar.jsx";
import Title from "../components/recommandMarker/title.jsx";
import Logo from "../components/recommandMarker/logo.jsx";
import PinMarker from "../components/addWating/PinMarker.jsx";
import ScrollBar from "../components/ScrollBar.jsx";
import ScrollBarBig from "../components/ScrollBarBig.jsx";
import Marker from "../components/recommandMarker/Marker.jsx";
import Button from "../components/pins/Button.jsx";
import ButtonBig from "../components/pins/ButtonBig.jsx";
import PinComponent from "../components/pins/PinComponent.jsx";
import hospitalIcon from "../components/icons/hospital.svg";

const Main = () => {
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
        <Logo></Logo>
        <button
          style={{
            backgroundColor: "#FF772A",
            borderRadius: "30px",
            width: "80vw",
            left: "10%",
            height: "5vh",
            position: "absolute",
            top: "230px",
            border: "none",
          }}
        >
          <p style={{ color: "white", fontWeight: "bold" }}>
            + 내 대기번호 추가
          </p>
        </button>
        <div
          style={{
            display: "flex",
            position: "absolute",
            padding: "0px",
            top: "1000%",
            left: "8%",
          }}
        >
          <img src={hospitalIcon} alt="Hospital Icon" />
          <p style={{ fontWeight: "bold", fontSize: "20px", color: "#FF772A" }}>
            유성구
          </p>
          <p style={{ fontWeight: "bold", fontSize: "20px" }}>
            {" "}
            내 다른 소아과
          </p>
        </div>
        <Button></Button>
        <PinComponent></PinComponent>
      </div>
    </div>
  );
};

Main.defaultProps = {
  name: "서현",
};

export default Main;
