import React from "react";
import { useState } from "react";
import { useEffect } from "react";
import MediaQuery from "react-responsive";
import ComponentsPinScroll from "./ComponentsPinScroll";

const PinComponent = (props) => {
  const { value, location, locationAccount, heart, walk } = props;
  return (
    <div>
      <div style={{ position: "fixed", bottom: "55%", left: "10%" }}>
        <ComponentsPinScroll
          value={value}
          location={"유성병원"}
          locationAccount={"북유성대로 93 · 1588-7011"}
          heart={22}
          walk={33}
        ></ComponentsPinScroll>
      </div>
      <div style={{ position: "fixed", bottom: "40%", left: "10%" }}>
        <ComponentsPinScroll
          value={value}
          location={"한사랑의원"}
          locationAccount={"봉명동 계룡로 150 "}
          heart={20}
          walk={30}
        ></ComponentsPinScroll>
      </div>
      <div style={{ position: "fixed", bottom: "25%", left: "10%" }}>
        <ComponentsPinScroll
          value={value}
          location={"튼튼병원"}
          locationAccount={"봉명동 559-4"}
          heart={10}
          walk={9}
        ></ComponentsPinScroll>
      </div>
    </div>
  );
};

export default PinComponent;
