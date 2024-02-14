import React from "react";
import MediaQuery from "react-responsive";
import pinMarkerIcon from "../icons/Frame 15.png";
import pinMarkerIconBig from "../icons/Frame 16.png";

const AddPin = () => {
  return (
    <div>
      <MediaQuery maxWidth={430}>
        {/* 아이폰 화면 크기에 맞는 스타일 */}
        <div style={{ position: "absolute", left: "30px", top: "210px" }}>
          <button
            style={{
              background: "white",
              border: "white",
            }}
          >
            <img src={pinMarkerIcon} alt="Pin Marker Icon" />
          </button>
        </div>
      </MediaQuery>

      <MediaQuery minWidth={431}>
        {/* 아이폰 화면 크기보다 큰 화면에 대한 스타일 */}
        <div>
          <div style={{ position: "absolute", left: "235px", top: "105px" }}>
            <button
              style={{
                background: "white",
                border: "white",
              }}
            >
              <img
                src={pinMarkerIconBig}
                alt="Pin Marker Icon Big"
                style={{ width: "auto", height: "15px" }}
              />
            </button>
          </div>
        </div>
      </MediaQuery>
    </div>
  );
};

export default AddPin;
