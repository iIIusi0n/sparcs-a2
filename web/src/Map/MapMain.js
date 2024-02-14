import React, { useEffect } from "react";
import MapboxSmall from "../components/MapboxSmall";

function Map() {
  return (
    <div style={{ position: "relative", top: "0", bottom: "0", width: "100%" }}>
      <MapboxSmall></MapboxSmall>
    </div>
  );
}

export default Map;
