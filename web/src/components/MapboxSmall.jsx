import React, { useEffect, useState } from "react";
import mapboxgl from "mapbox-gl";

function MapboxSmall() {
  const [markers, setMarkers] = useState([
    {
      longtitude: 127.3845475,
      latitude: 36.3504119,
      color: "#ffffff"
    },
  ]);

  useEffect(() => {
    mapboxgl.accessToken =
      "pk.eyJ1Ijoic29uZ2JvbmdqdW5lIiwiYSI6ImNsc2xmcjlybTA2dWkycm96cHlsanYyYXMifQ.jWx37L87Krg1pZ9NKm3mTQ"; // 여기에 Mapbox 액세스 토큰을 넣으세요
    const map = new mapboxgl.Map({
      container: "map",
      style: "mapbox://styles/mapbox/streets-v11",
      center: [127.3845475, 36.3504119], // 초기 위치 (경도, 위도)
      zoom: 14, // 초기 줌 레벨
    });

    /* const marker = new mapboxgl.Marker({
      color: "#FF0000", // 마커 색상
      draggable: false, // 사용자가 마커를 드래그할 수 있는지 여부
    })
      .setLngLat([127.3845475, 36.3504119]) // 마커의 위치 (경도, 위도)
      .addTo(map); */
    return () => {
      map.remove();
    }; // 컴포넌트가 언마운트될 때 맵을 제거합니다.
  }, []);

  return <div id="map" style={{ width: "100%", height: "1200px" }} />;
}

export default MapboxSmall;
