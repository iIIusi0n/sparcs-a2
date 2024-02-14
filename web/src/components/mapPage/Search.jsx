import React from "react";

function RoundedSearchEngine({ value, onChange }) {
  return (
    <div
      style={{
        display: "flex",
        alignItems: "center",
        borderRadius: "20px",
        padding: "15px 10px",
        width: "80vw",
        height: "3vh",
      }}
    >
      <input
        type="text"
        placeholder="검색어를 입력하세요..."
        value={value}
        onChange={onChange}
        style={{
          flex: "1",
          border: "none",
          outline: "none",
          padding: "5px",
          paddingLeft: "30px",
          borderRadius: "20px",
          height: "3vh",
        }}
      />
    </div>
  );
}

export default RoundedSearchEngine;
