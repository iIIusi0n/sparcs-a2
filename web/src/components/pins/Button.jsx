import React from "react";
import MediaQuery from "react-responsive";

const Button = () => {
  return (
    <div style={{ position: "relative", top: "320px", left: "20px" }}>
      <button
        style={{
          backgroundColor: "#FF772A",
          color: "white",
          position: "absolute",
          border: "none",
          borderRadius: "5px",
          left: "0",
        }}
      >
        대기 적은 순
      </button>
      <button
        style={{
          position: "absolute",
          left: "25%",
          backgroundColor: "#FFE8E1",
          color: "#FF772A",
          position: "absolute",
          border: "none",
          borderRadius: "5px",
        }}
      >
        거리 순
      </button>
      <button
        style={{
          position: "absolute",
          left: "43%",
          backgroundColor: "#FFE8E1",
          color: "#FF772A",
          position: "absolute",
          border: "none",
          borderRadius: "5px",
        }}
      >
        북 마크
      </button>
    </div>
  );
};

export default Button;
