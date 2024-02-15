import React from "react";
import backgroundIcon from "../icons/background.svg";
import bookMarkIcon from "../icons/Bookmark.svg";
import { useNavigate } from "react-router-dom";

const MessageMove = (props = "bongjune") => {
  const navigate = useNavigate();
  const handleClickToChatRoom = () => {
    // 버튼을 클릭하면 '/about' 경로로 이동
    navigate("/chatroom");
  };
  return (
    <div
      style={{
        display: "flex",
        width: "80vw",
        height: "10vh",
        marginRight: "10px",
        backgroundColor: "blue",
        borderRadius: "30px",
      }}
    >
      <div
        style={{
          borderRadius: "30px",
          position: "relative",
        }}
      >
        <img src={backgroundIcon} alt="이미지" />
        <div
          style={{
            position: "absolute",
            top: "15%",
            left: "7%",
            padding: "2px",
            color: "white",
            background: "transalate",
            borderRadius: "5px",
            fontSize: "15px",
            fontWeight: "normal",
          }}
        >
          대전광역시 유성구 전체
        </div>
        <div
          style={{
            position: "absolute",
            top: "40%",
            left: "7.5%",
            padding: "2px",
            color: "white",
            background: "transalate",
            borderRadius: "5px",
            fontSize: "10px",
            fontWeight: "normal",
          }}
        >
          운영자 | 1130명 참여 중
        </div>
        <div style={{ position: "absolute", left: "90%", top: "20%" }}>
          <img src={bookMarkIcon} alt="이미지" />
        </div>
        <button
          style={{
            backgroundColor: "#FF772A",
            borderRadius: "10px",
            width: "20vw",
            left: "70%",
            height: "3vh",
            position: "absolute",
            top: "150%",
            border: "none",
            fontSize: "10px",
            color: "white",
            fontWeight: "bold   ",
          }}
          onClick={handleClickToChatRoom}
        >
          참여하기
        </button>
      </div>
    </div>
  );
};

export default MessageMove;
