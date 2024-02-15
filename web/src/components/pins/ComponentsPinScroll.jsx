import React from "react";
import walkIcon from "../icons/walk.png";
import heartIcon from "../icons/heart.png";

const ComponentsPinScroll = ({
  value,
  location = "장소 이름이 들어갑니다",
  locationAccount = "장소에 대한 설명이 이 자리에 들어갑니다",
  heart = "22",
  walk = "22",
}) => {
  return (
    <div style={{ position: "relative", left: "15px", top: "360px" }}>
      <div
        style={{
          width: "130px",
          height: "100px",
          borderRadius: "10px",
          backgroundColor: "#A3A5A8",
          overflow: "hidden",
          position: "absolute",
        }}
      >
        <button>
          <img src={value} alt="이미지" />
        </button>
      </div>
      <div style={{ position: "absolute", left: "150px" }}>
        <p
          style={{
            fontSize: "18px",
            color: "#FF772A",
            fontWeight: "bold",
            fontFamily: "Pretendard-Bold",
          }}
        >
          {location}
        </p>
      </div>
      <div style={{ position: "absolute", left: "150px", top: "25px" }}>
        <p
          style={{
            fontSize: "14px",
            fontWeight: "bold",
            fontFamily: "Pretendard-Bold",
          }}
        >
          {locationAccount}
        </p>
      </div>
      <div
        style={{
          position: "absolute",
          left: "155px",
          top: "70px",
          display: "flex",
        }}
      >
        <img src={heartIcon} alt="Heart Icon" style={{ marginRight: "5px" }} />
        <p
          style={{
            fontSize: "12px",
            margin: "0px",
            fontFamily: "Pretendard-Medium",
          }}
        >
          {heart}
        </p>
        <img
          src={walkIcon}
          alt="Walk Icon"
          style={{ marginLeft: "30px", marginRight: "5px" }}
        />
        <p
          style={{
            fontSize: "12px",
            margin: "0px",
            fontFamily: "Pretendard-Medium",
          }}
        >
          {walk}
        </p>
      </div>
    </div>
  );
};

export default ComponentsPinScroll;
