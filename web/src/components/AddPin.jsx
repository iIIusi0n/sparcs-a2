import React from "react";
import MediaQuery from "react-responsive";
import addIcon from "../components/icons/add.png";

const AddPin = () => {
  return (
    <div>
      <MediaQuery maxWidth={430}>
        {/* 아이폰 화면 크기에 맞는 스타일 */}
        <div style={{ position: "absolute", left: "340px", top: "140px" }}>
          <img src={addIcon} alt="Add Icon" />
        </div>
      </MediaQuery>

      <MediaQuery minWidth={431}>
        {/* 아이폰 화면 크기보다 큰 화면에 대한 스타일 */}
        <div>
          <div style={{ position: "absolute", left: "390px", top: "55px" }}>
            <img
              src={addIcon}
              alt="Add Icon"
              style={{ width: "auto", height: "40px" }}
            />
          </div>
        </div>
      </MediaQuery>
    </div>
  );
};

export default AddPin;
