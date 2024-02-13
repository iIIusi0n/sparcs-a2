import React, { useState } from "react";
import styled from "styled-components";
import Header from "./components/header";
import NavigationBar from "../components/NavigationBar";
import MediaQuery from "react-responsive";
import Title from "./components/title";
import AddPin from "../components/AddPin.jsx";
import PinMarker from "../components/PinMarker.jsx";
import ScrollBar from "../components/ScrollBar.jsx";
import Marker from "../components/Marker.jsx";
const Container = styled.div`
  display: flex;
  flex-direction: column;
  padding-left: 10px;
`;

const Main = (props) => {
  return (
    <div>
      <MediaQuery maxWidth={430}>
        {/* 아이폰 화면 크기에 맞는 스타일 */}
        <div>
          <Container>
            <Header></Header>
            <Title></Title>
            <AddPin></AddPin>
            <PinMarker></PinMarker>
            <ScrollBar></ScrollBar>
            <Marker></Marker>
            <NavigationBar></NavigationBar>
          </Container>
        </div>
      </MediaQuery>

      <MediaQuery minWidth={431}>
        {/* 아이폰 화면 크기보다 큰 화면에 대한 스타일 */}
        <div>
          <Container>
            <Header></Header>
            <Title></Title>
            <AddPin></AddPin>
            <PinMarker></PinMarker>
            <NavigationBar></NavigationBar>
          </Container>
        </div>
      </MediaQuery>
    </div>
  );
};

Main.defaultProps = {
  name: "서현",
};

export default Main;
