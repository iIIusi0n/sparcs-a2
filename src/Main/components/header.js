import React, { useState } from "react";
import styled from "styled-components";
import { FaMapMarker } from "react-icons/fa";
import { FaSearch } from "react-icons/fa";
import { HiOutlineBellAlert } from "react-icons/hi2";

const Avatar = styled.div`
  position: relative;
  top: 20px;
  left: 26px;
`;
const Search = styled.div`
  position: relative;
  top: 20px;
  left: 100px;
`;
const Alert = styled.div`
  position: relative;
  top: 20px;
  left: 100px;
`;
const Header = () => {
  const [currentValue, setCurrentValue] = useState("대전광역시 동구");

  const map_index = [
    "대전광역시 서구",
    "대전광역시 중구",
    "대전광역시 유성구",
    "대전광역시 동구",
    "대전광역시 대덕구",
  ];
  const handleChange = (event) => {
    setCurrentValue(event.target.value);
  };

  return (
    <div style={{ display: "flex" }}>
      <Avatar>
        <FaMapMarker size={15} color="FF772A" />
      </Avatar>
      <select
        onChange={handleChange}
        style={{
          border: "none",
          outline: "none",
          position: "relative",
          top: "20px",
          left: "31px",
          fontSize: "8px",
          fontFamily: "Pretendard",
        }}
        value={currentValue}
      >
        <option value={map_index[0]}>{map_index[0]}</option>
        <option value={map_index[1]}>{map_index[1]}</option>
        <option value={map_index[2]}>{map_index[2]}</option>
        <option value={map_index[3]}>{map_index[3]}</option>
        <option value={map_index[4]}>{map_index[4]}</option>
      </select>
      <Search>
        <FaSearch size={15} color="FF772A" />
      </Search>
      <Alert>
        <HiOutlineBellAlert size={15} color="FF772A" />
      </Alert>
    </div>
  );
};

export default Header;
