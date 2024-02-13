import React from "react";
import url from "./icons/Rectangle.png";

const Container = (props) => {
  return (
    <div
      style={{
        display: "flex",
        height: "100vh",
      }}
    >
      <div style={{ borderRadius: "10px", overflow: "hidden" }}>
        <img src={props.url} alt="이미지" />
      </div>
    </div>
  );
};

Container.defaultProps = {
  url: url,
};

export default Container;
