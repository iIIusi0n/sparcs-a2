import React, { useEffect } from "react";
import MapboxSmall from "../components/mapPage/MapboxSmall";
import { useState } from "react";
import Search from "../components/mapPage/Search";
import searchIcon from "../components/icons/search.svg";
import currentButtonIcon from "../components/icons/currentbutton.svg";
import writeButtonIcon from "../components/icons/writebutton.svg";
import NavigationBarMap from "../components/navigation/NavigationBarMap";

function MapMain() {
  const [searchTerm, setSearchTerm] = useState("");

  const handleSearchChange = (e) => {
    setSearchTerm(e.target.value);
  };

  const handleSearch = () => {
    console.log("검색어:", searchTerm);
  };
  return (
    <div>
      <div
        style={{ position: "relative", top: "0", bottom: "0", width: "100%" }}
      >
        <div style={{ postion: "relative" }}>
          <MapboxSmall></MapboxSmall>
        </div>
        <div
          style={{
            position: "absolute",
            top: "2%",
            left: "8%",
            display: "flex",
          }}
        >
          <div style={{ position: "absolute", top: "23%", left: "1%" }}>
            <img
              src={searchIcon}
              alt="Search Icon"
              style={{ position: "absolute", width: "23px", left: "13px" }}
            />
          </div>
          <Search value={searchTerm} onChange={handleSearchChange}></Search>
          <button
            style={{
              position: "relative",
              top: "630px",
              left: "-390px",
              margin: "0px",
              padding: "0px",
              border: "0px",
            }}
          >
            <img src={currentButtonIcon} alt="Current Button Icon" />
          </button>
          <button
            style={{
              position: "relative",
              top: "630px",
              left: "-90px",
              margin: "0px",
              padding: "0px",
              border: "0px",
            }}
          >
            <img
              src={writeButtonIcon}
              alt="Write Button Icon"
              style={{ width: "72" }}
            />
          </button>
        </div>
        <div style={{ position: "absolute", bottom: "50%" }}>
          <NavigationBarMap></NavigationBarMap>
        </div>
        <div>
          <NavigationBarMap></NavigationBarMap>
        </div>
      </div>
    </div>
  );
}

export default MapMain;
