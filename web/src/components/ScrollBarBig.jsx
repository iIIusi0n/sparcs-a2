import React from "react";
import { useEffect } from "react";
import { useState } from "react";
import Container from "./Container";

const ScrollBarBig = ({ newValue, newLocation, newDate }) => {
  const [items, setItems] = useState([]);

  useEffect(() => {
    // 새로운 값을 배열에 추가
    setItems((prevItems) => [...prevItems, newValue]);
  }, [newValue, newLocation, newDate]); // newValue가 변경될 때만 이 효과 실행
  return (
    <div
      style={{
        width: "29%",
        overflowX: "auto",
        position: "relative",
        left: "230px",
        top: "110px",
        height: "auto",
        backgroundColor: "rgba(255, 255, 255, 0.5)",
      }}
    >
      <div style={{ display: "flex" }}>
        {items.map((item, index) => (
          <button
            style={{
              backgroundColor: "rgba(255, 255, 255, 0.5)",
              border: "white",
              marginColor: "white",
              background: "white",
            }}
          >
            <Container value={item}></Container>
          </button>
        ))}
      </div>
    </div>
  );
};

export default ScrollBarBig;
