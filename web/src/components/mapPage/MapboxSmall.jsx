import React, { useEffect, useState } from "react";
import mapboxgl from "mapbox-gl";
import ReactMapGL from "react-map-gl";
import "mapbox-gl/dist/mapbox-gl.css";
import mapMarkerIcon from "../icons/marker.svg";

function MapboxSmall() {
  const [markers, setMarkers] = useState([
    { longtitude: 127.3845475, latitude: 36.3504119, color: "#FF0000" },
  ]);

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
      el.style.backgroundColor = marker.color;
      el.style.backgroundImage = "url(${mapMarkerIcon})";
      el.style.width = "20px";
      el.style.height = "20px";

      new mapboxgl.Marker(el)
        .setLngLat([marker.longtitude, marker.latitude])
        .setDraggable(false)
        .addTo(map);
    });
    return () => {
      map.remove();
    }; // 컴포넌트가 언마운트될 때 맵을 제거합니다.
  }, []);

  return <div id="map" style={{ width: "100vw", height: "100vh" }} />;
}

export default MapboxSmall;
