import React from "react";
import url from "./icons/Rectangle.png";

const Container = ({
  value = url,
  location = "동구",
  date = " 2023년 12월 ",
}) => {
  return (
    <div
      style={{
        display: "flex",
        height: "100vh",
        marginRight: "10px",
      }}
    >
      <div
        style={{
          borderRadius: "10px",
          overflow: "hidden",
          position: "relative",
        }}
      >
        <img src={value} alt="이미지" />
        <div
          style={{
            position: "absolute",
            top: "7px",
            left: "7px",
            padding: "2px",
            color: "white",
            backgroundColor: "rgba(0, 0, 0, 0.5)",
            borderRadius: "5px",
            fontSize: "10px",
          }}
        >
          {location}
        </div>
        <div
          style={{
            position: "absolute",
            top: "200px",
            left: "60px",
            padding: "2px",
            color: "white",
            backgroundColor: "rgba(0, 0, 0, 0.5)",
            borderRadius: "5px",
            fontSize: "10px",
          }}
        >
          {date}
        </div>
      </div>
    </div>
  );
};

export default Container;
