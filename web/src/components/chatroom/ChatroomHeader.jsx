import React from "react";
import { FaSearch } from "react-icons/fa";
import menuIcon from "../icons/menu.svg";

const header = () => {
  return (
    <div>
      {/* 아이폰 화면 크기에 맞는 스타일 */}
      <div style={{ display: "flex", position: "relative" }}>
        <search style={{ position: "absolute", top: "10px", right: "15%" }}>
          <button
            style={{
              background: "white",
              border: "white",
              margin: "0px",
            }}
          >
            <FaSearch size={25} color="FF772A" />
          </button>
        </search>
        <alert
          style={{
            position: "absolute",
            top: "50%",
            right: "5%",
          }}
        >
          <button
            style={{
              background: "white",
              border: "white",
              margin: "0px",
              position: "absolute",
              top: "8px",
              right: "10%",
            }}
          >
            <img src={menuIcon} alt="Profile Icon" />
          </button>
        </alert>
      </div>
    </div>
  );
};

export default header;
