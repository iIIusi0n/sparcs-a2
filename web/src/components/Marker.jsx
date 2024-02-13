import React from "react";
import MediaQuery from "react-responsive";
import markerIcon from "./icons/Group.png";

const Marker = () => {
  return (
    <div>
      <MediaQuery maxWidth={430}>
        {/* 아이폰 화면 크기에 맞는 스타일 */}
        <div
          style={{
            position: "absolute",
            left: "25px",
            top: "540px",
            display: "flex",
          }}
        >
          <img src={markerIcon} alt="Marker Icon" />
          <p
            style={{
              position: "relative",
              top: "2px",
              left: "10px",
              fontWeight: "bold",
              margin: "0px",
              border: "0px",
              fontSize: "20px",
            }}
          >
            지금 떠오르는 PIN
          </p>
        </div>
      </MediaQuery>

      <MediaQuery minWidth={431}>
        {/* 아이폰 화면 크기보다 큰 화면에 대한 스타일 */}
      </MediaQuery>
    </div>
  );
};

export default Marker;
