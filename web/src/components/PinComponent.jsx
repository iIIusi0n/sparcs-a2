import React from "react";
import Container from "./Container";
import url from "./icons/Rectangle.png";
import MediaQuery from "react-responsive";

const PinComponent = ({
  value = url,
  location = "동구",
  date = " 2023년 12월 ",
}) => {
  return (
    <div>
      <MediaQuery maxWidth={430}>
        <div
          style={{
            position: "relative",
            left: "100px",
            top: "350px",
            width: "100px",
          }}
        ></div>
      </MediaQuery>
      <MediaQuery maxWidth={431}>
        <div></div>
      </MediaQuery>
    </div>
  );
};

export default PinComponent;
