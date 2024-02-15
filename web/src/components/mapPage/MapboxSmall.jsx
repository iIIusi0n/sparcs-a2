import React, { useEffect, useState } from "react";
import mapboxgl from "mapbox-gl";
import ReactMapGL from "react-map-gl";
import "mapbox-gl/dist/mapbox-gl.css";
import mapMarkerIcon_orange from "../icons/marker.svg";
import mapMarkerIcon_red from "../icons/red_marker.svg";
import mapMarkerIcon_green from "../icons/green_marker.svg";
import selectMarkerIcon from "../icons/selectmarker.svg";
import currentButtonIcon from "../icons/currentbutton.svg";
import writeButtonIcon from "../icons/writebutton.svg";

function MapboxSmall() {
  const [markers, setMarkers] = useState([
    { longtitude: 127.3845475, latitude: 36.3504119, number: 22 },
  ]);

  const [selectedMarker, setSelectedMarker] = useState(null);
  useEffect(() => {
    mapboxgl.accessToken =
      "pk.eyJ1Ijoic29uZ2JvbmdqdW5lIiwiYSI6ImNsc2xmcjlybTA2dWkycm96cHlsanYyYXMifQ.jWx37L87Krg1pZ9NKm3mTQ"; // 여기에 Mapbox 액세스 토큰을 넣으세요
    const map = new mapboxgl.Map({
      container: "map",
      style: "mapbox://styles/mapbox/streets-v11",
      center: [127.3845475, 36.3504119], // 초기 위치 (경도, 위도)
      attributionControl: false,
      zoom: 14, // 초기 줌 레벨
    });

    markers.forEach((marker) => {
      const el = document.createElement("div");
      el.className = "marker";
      el.style.backgroundImage = getMarkerColor(marker.number);
      el.style.width = "50px";
      el.style.height = "55px";

      new mapboxgl.Marker(el)
        .setLngLat([marker.longtitude, marker.latitude])
        .setDraggable(false)
        .addTo(map);

      el.addEventListener("click", () => {
        el.style.backgroundImage = `url(${selectMarkerIcon})`;
        handleMarkerClick(marker);
      });
    });
    return () => {
      map.remove();
    }; // 컴포넌트가 언마운트될 때 맵을 제거합니다.
  }, []);
  const getMarkerColor = (number) => {
    if (number <= 10) {
      return `url(${mapMarkerIcon_green})`;
    } else if (number <= 20) {
      return `url(${mapMarkerIcon_orange})`;
    } else {
      return `url(${mapMarkerIcon_red})`;
    }
  };
  const handleMarkerClick = (marker) => {
    setSelectedMarker(marker);
    // 여기에서 클릭한 마커에 대한 추가 작업을 수행할 수 있습니다.
  };
  return (
    <div>
      <div id="map" style={{ width: "100vw", height: "100vh" }}></div>
      {selectedMarker && (
        <div>
          <div style={{ postion: "relative" }}>
            <button
              style={{
                position: "absolute",
                top: "630px",
                left: "10px",
                margin: "0px",
                padding: "0px",
                border: "0px",
              }}
            >
              <img src={currentButtonIcon} alt="Current Button Icon" />
            </button>
            <button
              style={{
                position: "absolute",
                top: "615px",
                right: "5px",
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
          <div
            style={{
              backgroundColor: "white",
              width: "100vw",
              position: "absolute",
              bottom: "55px",
            }}
          >
            <h2>Marker Info</h2>
            <p>Longitude: {selectedMarker.longitude}</p>
            <p>Latitude: {selectedMarker.latitude}</p>
            <p>Number: {selectedMarker.number}</p>
          </div>
        </div>
      )}
    </div>
  );
}

export default MapboxSmall;
