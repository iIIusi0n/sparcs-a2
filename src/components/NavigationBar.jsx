import React from "react";
import homeIcon from "../components/icons/Vector (Stroke).svg";
import styled from "styled-components";

const Container = styled.div`
  position: fixed;
  bottom: 5px;
  left: 50%;
  width: 100%;
`;

const Font = styled.div`
  position: relative;
  margin: 0%;
  font-size: 2%,
  font-weight: bold,
  text-align: center
`;

const NavigationBar = () => {
  return (
    <Container>
      <div>
        <img src={homeIcon} alt="Home Icon" />
        <Font>홈</Font>
      </div>
    </Container>
  );
};

export default NavigationBar;
