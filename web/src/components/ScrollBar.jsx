import React from "react";
import { useEffect } from "react";
import MediaQuery from "react-responsive";
import url from "./icons/Rectangle.png";
import Container from "./Container";

const ScrollBar = (props) => {
  const newImage = [];

  return (
    <div>
      <MediaQuery maxWidth={430}>
        <div
          style={{
            width: "50%",
            overflowX: "auto",
            position: "relative",
            top: "260px",
          }}
        >
          <div style={{ display: "flex" }}>
            <Container props={newImage}></Container>
          </div>
        </div>
      </MediaQuery>
      <MediaQuery maxWidth={431}></MediaQuery>
    </div>
  );
};

ScrollBar.defaultProps = {
  image: url,
};

export default ScrollBar;
