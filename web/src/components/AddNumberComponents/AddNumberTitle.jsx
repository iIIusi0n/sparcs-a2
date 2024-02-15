import React from "react";
import { useState, useEffect } from "react";
import NavigationBar from "../navigation/NavigationBar";
import iconLogo from "../icons/iconlogo.svg";
import 

const AddNumberTitle = (props) => {
  return (
    <div>
      {/* 아이폰 화면 크기에 맞는 스타일 */}
      <div>
        <div style={{ position: "absolute", top: "15%", left: "10%" }}>
          <div>
            <p
              style={{
                fontSize: "20px",
                fontFamily: "pretendard-Bold",
                margin: "0px",
              }}
            >
              접수를 완료하셨나요?
            </p>
          </div>
          <div style={{ display: "flex" }}>
            <p
              style={{
                fontFamily: "pretendard-Bold",
                fontSize: "20px",
                marginTop: "10px",
              }}
            >
              {props.userId}님의{" "}
            </p>
            <p
              style={{
                fontFamily: "pretendard-Bold",
                color: "#FF772A",
                fontSize: "20px",
                marginLeft: "8px",
                marginTop: "10px",
              }}
            >
              대기번호
            </p>
            <p
              style={{
                fontFamily: "pretendard-Bold",
                fontSize: "20px",
                marginTop: "10px",
              }}
            >
              입력해주세요!
            </p>
          </div>
          <p
            style={{
              fontSize: "12px",
              fontFamily: "pretendard-Medium",
              marginTop: "1%",
              marginBottom: "1%",
            }}
          >
            대기번호를 입력하면 같은 지역 내 보호자들에게
          </p>
          <p
            style={{
              fontSize: "12px",
              fontFamily: "pretendard-Medium",
              marginTop: "0%",
            }}
          >
            실시간으로 대기 현황이 공유됩니다
          </p>
        </div>
        <div style={{ display: "flex" }}>
          <input
            style={{
              position: "absolute",
              width: "16vw",
              height: "5vh",
              backgroundColor: "#d1d2d3",
              left: "10%",
              top: "30%",
              borderRadius: "15px",
              border: "none",
            }}
          ></input>
          <p
            style={{
              position: "absolute",
              left: "28%",
              top: "26%",
              fontFamily: "pretendard-Bold",
              fontSize: "40px",
            }}
          >
            번
          </p>
        </div>

      </div>
      <NavigationBar></NavigationBar>
    </div>
  );
};

export default AddNumberTitle;
