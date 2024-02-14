import React, { useState } from "react";
import styled from "styled-components";
import MediaQuery from "react-responsive";

const Title = (props) => {
  return (
    <div>
      <MediaQuery maxWidth={430}>
        {/* 아이폰 화면 크기에 맞는 스타일 */}
        <div>
          <div style={{ position: "absolute", top: "50px", left: "40px" }}>
            <div>
              <p style={{ fontWeight: "bold", fontSize: "20px" }}>
                {props.name}님,
              </p>
              <p style={{ fontWeight: "bold", fontSize: "20px" }}>
                오늘 방문한 장소를
              </p>
              <div style={{ display: "flex", margin: "0px", border: "0px" }}>
                <p
                  style={{
                    fontWeight: "bold",
                    color: "#FF772A",
                    margin: "0px",
                    border: "0px",
                    fontSize: "20px",
                  }}
                >
                  PIN
                </p>
                <p
                  style={{
                    fontWeight: "bold",
                    margin: "0px",
                    border: "0px",
                    fontSize: "20px",
                  }}
                >
                  으로 남겨보세요
                </p>
              </div>
            </div>
          </div>
        </div>
      </MediaQuery>

      <MediaQuery minWidth={431}>
        {/* 아이폰 화면 크기보다 큰 화면에 대한 스타일 */}
        <div>
          <div style={{ position: "absolute", top: "40px", left: "250px" }}>
            <div>
              <p style={{ fontWeight: "bold", fontSize: "8px" }}>
                {props.name}님,
              </p>
              <p style={{ fontWeight: "bold", fontSize: "8px" }}>
                오늘 방문한 장소를
              </p>
              <div
                style={{
                  display: "flex",
                  margin: "0px",
                  border: "0px",
                  fontSize: "8px",
                }}
              >
                <p
                  style={{
                    fontWeight: "bold",
                    color: "#FF772A",
                    margin: "0px",
                    border: "0px",
                  }}
                >
                  PIN
                </p>
                <p style={{ fontWeight: "bold", margin: "0px", border: "0px" }}>
                  으로 남겨보세요
                </p>
              </div>
            </div>
          </div>
        </div>
      </MediaQuery>
    </div>
  );
};

Title.defaultProps = {
  name: "서현",
};

export default Title;
