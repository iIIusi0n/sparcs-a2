import React from "react";
import { useState, useEffect } from "react";
import coffeIcon from "../icons/coffe.svg";

const Title = () => {
  return (
    <div>
      {/* 아이폰 화면 크기에 맞는 스타일 */}
      <div>
        <div style={{ position: "absolute", top: "200%", left: "40px" }}>
          <div>
            <p style={{ fontWeight: "bold", fontSize: "20px" }}>
              지루한 진료 대기 시간,
            </p>
            <div style={{ display: "flex", margin: "0px", border: "0px" }}>
              <p
                style={{
                  fontWeight: "bold",
                  fontSize: "20px",
                  margin: "0px",
                  border: "0px",
                }}
              >
                동네 부모님들끼리
              </p>
              <p
                style={{
                  fontWeight: "bold",
                  color: "#FF772A",
                  margin: "0px",
                  border: "0px",
                  fontSize: "20px",
                }}
              >
                소통해요
              </p>
              <img src={coffeIcon} alt="Coffe Icon" />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

Title.defaultProps = {
  name: "서현",
  time: "1시간",
  number: 18,
};

export default Title;
