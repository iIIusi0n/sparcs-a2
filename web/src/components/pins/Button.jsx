import React from "react";
import MediaQuery from "react-responsive";

const Button = () => {
  return (
    <div>
      <MediaQuery maxWidth={430}>
        <div style={{ position: "relative", top: "320px", left: "20px" }}>
          <button
            style={{
              backgroundColor: "#FF772A",
              color: "white",
              position: "absolute",
              border: "none",
              borderRadius: "5px",
            }}
          >
            맟춤 추천
          </button>
          <button
            style={{
              position: "absolute",
              left: "90px",
              backgroundColor: "#FFE8E1",
              color: "#FF772A",
              position: "absolute",
              border: "none",
              borderRadius: "5px",
            }}
          >
            20대 인기
          </button>
          <button
            style={{
              position: "absolute",
              left: "180px",
              backgroundColor: "#FFE8E1",
              color: "#FF772A",
              position: "absolute",
              border: "none",
              borderRadius: "5px",
            }}
          >
            좋아요 순
          </button>
          <button
            style={{
              position: "absolute",
              left: "270px",
              backgroundColor: "#FFE8E1",
              color: "#FF772A",
              position: "absolute",
              border: "none",
              borderRadius: "5px",
            }}
          >
            거리 순
          </button>
        </div>
      </MediaQuery>
    </div>
  );
};

export default Button;
