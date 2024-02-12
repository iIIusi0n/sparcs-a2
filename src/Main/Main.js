import React, { useState } from "react";
import styled from "styled-components";
import Header from "./components/header";
import NavigationBar from "../components/NavigationBar";

const Container = styled.div`
  display: flex;
  flex-direction: column;
  @media (max-width: 1284px) {
    /* 아이폰 15 Pro 화면 크기에 맞춘 스타일 */
    padding-left: 200px;
    max-width: 1284px;
    /* 원하는 스타일 */
  }
`;

const Main = () => {
  return (
    <Container>
      <Header></Header>
      <NavigationBar></NavigationBar>
    </Container>
  );
};
export default Main; //
