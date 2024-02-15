import React from "react";
import { useState, useEffect } from "react";

const Title = (props) => {
  const [currentTime, setCurrentTime] = useState(new Date());

  useEffect(() => {
    const timer = setInterval(() => {
      setCurrentTime(new Date());
    }, 1000);

    return () => {
      clearInterval(timer);
    };
  }, []);

  const formattedTime = currentTime.toLocaleTimeString([], {
    hour: "2-digit",
    minute: "2-digit",
  });
  return (
    <div>
      {/* 아이폰 화면 크기에 맞는 스타일 */}
      <div>
        <div style={{ position: "absolute", top: "50px", left: "40px" }}>
          <div>
            <p style={{ fontWeight: "bold", fontSize: "20px" }}>
              {props.name}님,
            </p>
            <p style={{ fontWeight: "bold", fontSize: "20px" }}>
              예상 대기 시간은
            </p>
            <div style={{ display: "flex", margin: "0px", border: "0px" }}>
              <p
                style={{
                  fontWeight: "bold",
                  fontSize: "20px",
                  margin: "0px",
                  border: "0px",
                  color: "#FF772A",
                }}
              >
                최소
              </p>
              <p
                style={{
                  fontWeight: "bold",
                  color: "#FF772A",
                  margin: "0px",
                  border: "0px",
                  fontSize: "20px",
                }}
              >
                {props.time}
              </p>
              <p
                style={{
                  fontWeight: "bold",
                  margin: "0px",
                  border: "0px",
                  fontSize: "20px",
                }}
              >
                이예요
              </p>
            </div>
          </div>
          <p style={{ fontSize: "12px" }}>
            {formattedTime}기준 대기번호 {props.number}번
          </p>
        </div>
      </div>
    </div>
  );
};

Title.defaultProps = {
  name: "서현",
  time: "1시간",
  number: 18,
};

export default Title;
