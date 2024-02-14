import React from "react";
import homeIcon from "../components/icons/Frame 14.png";
import pinIcon from "../components/icons/map.png";
import messageIcon from "../components/icons/ellipsis.bubble.png";
import profileIcon from "../components/icons/iconusercirclealt.png";
import styled from "styled-components";
import MediaQuery from "react-responsive";

const NavigationBar = () => {
  return (
    <div>
      <MediaQuery maxWidth={430}>
        {/* 아이폰 화면 크기에 맞는 스타일 */}
        <div
          style={{
            background: "white",
            width: "100%",
            border: "black",
          }}
        >
          <div
            style={{
              position: "fixed",
              bottom: "15px",
              left: "80px",
              display: "flex",
              flexDirection: "column",
            }}
          >
            <button
              style={{
                background: "white",
                border: "white",
                margin: "0px",
              }}
            >
              <img src={homeIcon} alt="Home Icon" />
            </button>
            <p
              style={{
                fontSize: "14px",
                fontFamily: "Pretendard",
                position: "relative",
                margin: "0",
                left: "14px",
                color: "#FF772A",
              }}
            >
              홈
            </p>
          </div>
          <div
            style={{
              position: "fixed",
              bottom: "15px",
              left: "160px",
              display: "flex",
              flexDirection: "column",
            }}
          >
            <button
              style={{
                background: "white",
                border: "white",
                margin: "0px",
              }}
            >
              <img src={pinIcon} alt="Pin Icon" />
            </button>
            <p
              style={{
                fontSize: "14px",
                fontFamily: "Pretendard",
                position: "relative",
                margin: "0",
                left: "8px",
              }}
            >
              PIN
            </p>
          </div>
          <div
            style={{
              position: "fixed",
              bottom: "15px",
              left: "240px",
              display: "flex",
              flexDirection: "column",
            }}
          >
            <button
              style={{
                background: "white",
                border: "white",
                margin: "0px",
              }}
            >
              <img src={messageIcon} alt="Message Icon" />
            </button>
            <p
              style={{
                fontSize: "14px",
                fontFamily: "Pretendard",
                position: "relative",
                left: "4px",
                margin: "0",
              }}
            >
              모임
            </p>
          </div>

          <div
            style={{
              position: "fixed",
              bottom: "17px",
              left: "320px",
              display: "flex",
              flexDirection: "column",
            }}
          >
            <button
              style={{
                background: "white",
                border: "white",
                margin: "0px",
              }}
            >
              <img src={profileIcon} alt="Profile Icon" />
            </button>
            <p
              style={{
                fontSize: "14px",
                fontFamily: "Pretendard",
                position: "relative",
                margin: "0",
                top: "4px",
                left: "2px",
              }}
            >
              마이
            </p>
          </div>
        </div>
      </MediaQuery>
      <MediaQuery minWidth={431}>
        {/* 아이폰 화면 크기보다 큰 화면에 대한 스타일 */}
        <div
          style={{
            background: "white",
            width: "100%",
            border: "black",
          }}
        >
          <div
            style={{
              position: "fixed",
              bottom: "15px",
              left: "245px",
              display: "flex",
              flexDirection: "column",
            }}
          >
            <button
              style={{
                backgroundColor: "rgba(255, 255, 255, 0.5)",
                border: "white",
                margin: "0px",
              }}
            >
              <img
                src={homeIcon}
                alt="Home Icon"
                style={{ width: "auto", height: "20px" }}
              />
            </button>
            <p
              style={{
                fontSize: "8px",
                fontFamily: "Pretendard",
                position: "relative",
                margin: "0",
                left: "12px",
                color: "#FF772A",
              }}
            >
              홈
            </p>
          </div>
          <div
            style={{
              position: "fixed",
              bottom: "15px",
              left: "290px",
              display: "flex",
              flexDirection: "column",
            }}
          >
            <button
              style={{
                backgroundColor: "rgba(255, 255, 255, 0.5)",
                border: "white",
                margin: "0px",
              }}
            >
              <img
                src={pinIcon}
                alt="Pin Icon"
                style={{ width: "auto", height: "20px" }}
              />
            </button>
            <p
              style={{
                fontSize: "8px",
                fontFamily: "Pretendard",
                position: "relative",
                margin: "0",
                left: "10px",
              }}
            >
              PIN
            </p>
          </div>
          <div
            style={{
              position: "fixed",
              bottom: "15px",
              left: "335px",
              display: "flex",
              flexDirection: "column",
            }}
          >
            <button
              style={{
                backgroundColor: "rgba(255, 255, 255, 0.5)",
                border: "white",
                margin: "0px",
              }}
            >
              <img
                src={messageIcon}
                alt="Message Icon"
                style={{ width: "auto", height: "20px" }}
              />
            </button>
            <p
              style={{
                fontSize: "8px",
                fontFamily: "Pretendard",
                position: "relative",
                left: "8px",
                margin: "0",
              }}
            >
              모임
            </p>
          </div>

          <div
            style={{
              position: "fixed",
              bottom: "15px",
              left: "380px",
              display: "flex",
              flexDirection: "column",
            }}
          >
            <button
              style={{
                backgroundColor: "rgba(255, 255, 255, 0.5)",
                border: "white",
                margin: "0px",
              }}
            >
              <img
                src={profileIcon}
                alt="Profile Icon"
                style={{ width: "auto", height: "20px" }}
              />
            </button>
            <p
              style={{
                fontSize: "8px",
                fontFamily: "Pretendard",
                position: "relative",
                margin: "0",
                top: "2px",
                left: "8px",
              }}
            >
              마이
            </p>
          </div>
        </div>
      </MediaQuery>
    </div>
  );
};

export default NavigationBar;
