import React from "react";
import { useState } from "react";
import { useEffect } from "react";
import MediaQuery from "react-responsive";
import ComponentsPinScroll from "./ComponentsPinScroll";

const PinComponent = (props) => {
  const { value, location, locationAccount, heart, walk } = props;
  return (
    <div>
      <div>
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
