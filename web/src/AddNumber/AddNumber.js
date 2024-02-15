import React from "react";
import backwardIcon from "../components/icons/backward.svg";
import AddNumberTitle from "../components/AddNumberComponents/AddNumberTitle";
import { useState, useEffect } from "react";

const AddNumber = () => {
  const [userId, setUserId] = useState("서현");

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
        <img src={backwardIcon} alt="Backward Icon" />
      </div>
      <AddNumberTitle userId={userId}></AddNumberTitle>
    </div>
  );
};

export default AddNumber;
