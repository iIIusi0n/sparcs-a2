import React from "react";
import MediaQuery from "react-responsive";

const ScrollBar = (props) => {
  return (
    <div>
      <MediaQuery maxWidth={430}>
        <div></div>
      </MediaQuery>
      <MediaQuery maxWidth={431}></MediaQuery>
    </div>
  );
};
