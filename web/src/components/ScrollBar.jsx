import React from "react";
import { useEffect } from "react";
import { useState } from "react";
import MediaQuery from "react-responsive";
import url from "./icons/Rectangle.png";
import Container from "./Container";

const ScrollBar = ({ newValue, newLocation, newDate }) => {
  const [items, setItems] = useState([]);

  useEffect(() => {
    // 새로운 값을 배열에 추가
    setItems((prevItems) => [...prevItems, newValue]);
  }, [newValue, newLocation, newDate]); // newValue가 변경될 때만 이 효과 실행
  return (
    <div>
      <MediaQuery maxWidth={430}>
        <div
          style={{
            width: "100%",
            overflowX: "auto",
            position: "relative",
            top: "260px",
          }}
        >
          <div style={{ display: "flex" }}>
            {items.map((item, index) => (
              <Container value={item}></Container>
            ))}
          </div>
        </div>
      </MediaQuery>
      <MediaQuery maxWidth={431}></MediaQuery>
    </div>
  );
};

ScrollBar.defaultProps = {
  image: url,
};

export default ScrollBar;
