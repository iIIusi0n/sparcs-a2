import React, { useEffect } from "react";
import MapboxSmall from "../components/mapPage/MapboxSmall";
import { useState } from "react";
import Search from "../components/mapPage/Search";
import searchIcon from "../components/icons/search.svg";

function MapMain() {
  const [searchTerm, setSearchTerm] = useState("");

  const handleSearchChange = (e) => {
    setSearchTerm(e.target.value);
  };

  const handleSearch = () => {
    console.log("검색어:", searchTerm);
  };
  return (
    <div style={{ position: "relative", top: "0", bottom: "0", width: "100%" }}>
      <div>
        <MapboxSmall style={{ postion: "absolute" }}></MapboxSmall>
      </div>
      <div
        style={{
          position: "absolute",
          top: "2%",
          left: "8%",
          display: "flex",
        }}
      >
        <div style={{ position: "absolute", top: "28%" }}>
          <img
            src={searchIcon}
            alt="Search Icon"
            style={{ position: "absolute", width: "23px", left: "13px" }}
          />
        </div>
        <Search value={searchTerm} onChange={handleSearchChange}></Search>
      </div>
    </div>
  );
}

export default MapMain;
