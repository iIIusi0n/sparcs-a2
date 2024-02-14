import React from "react";

const ButtonBig = () => {
  return (
    <div style={{ position: "relative", left: "235px" }}>
      <button
        style={{
          backgroundColor: "#FF772A",
          color: "white",
          position: "absolute",
          border: "none",
          borderRadius: "5px",
          fontSize: "8px",
        }}
      >
        맟춤 추천
      </button>
      <button
        style={{
          position: "absolute",
          left: "50px",
          backgroundColor: "#FFE8E1",
          color: "#FF772A",
          position: "absolute",
          border: "none",
          borderRadius: "5px",
          fontSize: "8px",
        }}
      >
        20대 인기
      </button>
      <button
        style={{
          position: "absolute",
          left: "100px",
          backgroundColor: "#FFE8E1",
          color: "#FF772A",
          position: "absolute",
          border: "none",
          borderRadius: "5px",
          fontSize: "8px",
        }}
      >
        좋아요 순
      </button>
      <button
        style={{
          position: "absolute",
          left: "150px",
          backgroundColor: "#FFE8E1",
          color: "#FF772A",
          position: "absolute",
          border: "none",
          borderRadius: "5px",
          fontSize: "8px",
        }}
      >
        거리 순
      </button>
    </div>
  );
};

export default ButtonBig;
