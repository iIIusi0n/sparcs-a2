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
          location={location}
          locationAccount={locationAccount}
          heart={heart}
          walk={walk}
        ></ComponentsPinScroll>
      </div>
      <div style={{ position: "fixed", bottom: "40%", left: "10%" }}>
        <ComponentsPinScroll
          value={value}
          location={location}
          locationAccount={locationAccount}
          heart={heart}
          walk={walk}
        ></ComponentsPinScroll>
      </div>
      <div style={{ position: "fixed", bottom: "25%", left: "10%" }}>
        <ComponentsPinScroll
          value={value}
          location={location}
          locationAccount={locationAccount}
          heart={heart}
          walk={walk}
        ></ComponentsPinScroll>
      </div>
    </div>
  );
};

export default PinComponent;
