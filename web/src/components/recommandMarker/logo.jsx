import React from "react";
import MediaQuery from "react-responsive";
import doctorIcon from "../icons/doctor.svg";

const AddPin = () => {
  return (
    <div>
      {/* 아이폰 화면 크기에 맞는 스타일 */}
      <div style={{ position: "absolute", left: "55%", top: "200%" }}>
        <img src={doctorIcon} alt="Doctor Icon" />
      </div>
    </div>
  );
};

export default AddPin;
