import React, { useState, useEffect } from "react";

function Location() {
  const [latitude, setLatitude] = useState(null);
  const [longitude, setLongitude] = useState(null);
  const [error, setError] = useState(null);

  const getLocation = () => {
    if ("geolocation" in navigator) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          setLatitude(position.coords.latitude);
          setLongitude(position.coords.longitude);
          setError(null);
        },
        (error) => {
          setError(error.message);
        }
      );
    } else {
      setError("Geolocation is not supported in this browser.");
    }
  };

  useEffect(() => {
    getLocation();
  }, []); // 컴포넌트가 마운트될 때 한 번만 실행

  return (
    <div>
      <button onClick={getLocation}>Get Location</button>
      {latitude && longitude && (
        <p>
          Latitude: {latitude}, Longitude: {longitude}
        </p>
      )}
      {error && <p>Error: {error}</p>}
    </div>
  );
}

export default Location;
