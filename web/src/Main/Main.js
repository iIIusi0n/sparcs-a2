import React, { useState } from "react";
import styled from "styled-components";
import Header from "../components/header/header.jsx";
import NavigationBar from "../components/navigation/NavigationBar.jsx";
import MediaQuery from "react-responsive";
import Title from "../components/recommandMarker/title.jsx";
import AddPin from "../components/recommandMarker/AddPin.jsx";
import PinMarker from "../components/addWating/PinMarker.jsx";
import ScrollBar from "../components/ScrollBar.jsx";
import ScrollBarBig from "../components/ScrollBarBig.jsx";
import Marker from "../components/recommandMarker/Marker.jsx";
import Button from "../components/pins/Button.jsx";
import ButtonBig from "../components/pins/ButtonBig.jsx";
import PinComponent from "../components/pins/PinComponent.jsx";
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
            <Button></Button>
            <PinComponent></PinComponent>
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
            <ScrollBarBig></ScrollBarBig>
            <Marker></Marker>
            <ButtonBig></ButtonBig>
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
