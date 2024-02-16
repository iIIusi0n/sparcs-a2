import React from "react";
import backwardIcon from "../components/icons/backward.svg";
import AddNumberTitle from "../components/AddNumberComponents/AddNumberTitle";
import { useState, useEffect } from "react";
import iconLogoIcon from "../components/icons/iconlogo.svg";
import { useNavigate } from "react-router-dom";

const AddNumber = () => {
  const navigate = useNavigate();
  const handleClickToHome = () => {
    // 버튼을 클릭하면 '/about' 경로로 이동
    navigate("/");
  };
  const [userId, setUserId] = useState("서현");
  const [inputValue, setInputValue] = useState();
  const handleInputChange = (event) => {
    setInputValue(event.target.value); // input 값이 변경될 때마다 상태를 업데이트합니다.
  };
  // 컴포넌트가 마운트될 때 한 번만 실행되는 useEffect
  useEffect(() => {
    // 서버로부터 아이디를 가져옵니다.
    fetch("/api/request-id")
      .then((response) => {
        if (response.ok) {
          return response.json();
        } else {
          throw new Error("Failed to request ID");
        }
      })
      .then((data) => {
        setUserId(data.id); // 받은 아이디를 상태에 설정합니다.
      })
      .catch((error) => {
        console.error("Error requesting ID:", error);
      });
  }, []); // 빈 배열을 전달하여 컴포넌트가 마운트될 때 한 번만 실행되도록 합니다.

  return (
    <div>
      <div style={{ position: "fixed", left: "7%", top: "3%" }}>
        <button
          style={{ backgroundColor: "white", border: "none" }}
          onClick={handleClickToHome}
        >
          <img src={backwardIcon} alt="Backward Icon" />
        </button>
      </div>
      <AddNumberTitle userId={userId}></AddNumberTitle>
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
          value={inputValue}
          onchange={handleInputChange}
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
      <img
        src={iconLogoIcon}
        alt="IconLogoIcon"
        style={{ position: "absolute", bottom: "15%", right: "0" }}
      />
      <button
        style={{
          width: "80vw",
          height: "7vh",
          position: "absolute",
          left: "10%",
          bottom: "8%",
          borderRadius: "10px",
          border: "none",
          color: "white",
          fontFamily: "pretendard-Bold",
          backgroundColor: "#FF772A",
          fontSize: "20px",
        }}
      >
        숫자를 입력해주세요
      </button>
    </div>
  );
};

export default AddNumber;
