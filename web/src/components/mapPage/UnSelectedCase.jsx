import React, { useEffect, useRef, useState } from "react";
import mapboxgl from "mapbox-gl";
import ReactMapGL from "react-map-gl";
import "mapbox-gl/dist/mapbox-gl.css";
import mapMarkerIcon_orange from "../icons/marker.svg";
import mapMarkerIcon_red from "../icons/red_marker.svg";
import mapMarkerIcon_green from "../icons/green_marker.svg";
import selectMarkerIcon_orange from "../icons/selectmarker.svg";
import selectMarkerIcon_red from "../icons/red_selectmarker.svg";
import selectMarkerIcon_green from "../icons/green_selectmarker.svg";
import currentButtonIcon from "../icons/currentbutton.svg";
import writeButtonIcon from "../icons/writebutton.svg";
import hospitalIcon from "../icons/hospital.svg";
import { useNavigate } from "react-router-dom";
import SelectedCase from "./SelectedCase";

function UnSelectedCase(props) {
  const getLocation = () => {
    if ("geolocation" in navigator) {
      navigator.geolocation.getCurrentPosition(
        function (position) {
          const latitude = position.coords.latitude;
          const longitude = position.coords.longitude;

          // 위치 정보를 서버로 전송
          fetch("/api/location", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({ latitude, longitude }),
          })
            .then((response) => {
              if (response.ok) {
                console.log("Location sent successfully");
              } else {
                console.error("Failed to send location");
              }
            })
            .catch((error) => {
              console.error("Error sending location:", error);
            });
        },
        function (error) {
          console.error("Failed to retrieve location:", error);
        }
      );
    } else {
      console.log("Geolocation is not supported by this browser.");
    }
  };

  const { name, location, etc, watingNum, watingTime, distance, image } = props;
  const [markers, setMarkers] = useState([
    { longitude: 127.3845475, latitude: 36.3505119, number: 22 },
    { longitude: 127.3822485, latitude: 36.3482102, number: 2 },
    { longitude: 127.3813485, latitude: 36.3493102, number: 13 },
    { longitude: 127.3843485, latitude: 36.3433102, number: 13 },
  ]);

  function usePrevious(value) {
    const ref = useRef();
    useEffect(() => {
      ref.current = value;
    });
    return ref.current;
  }

  const [selectedMarker, setSelectedMarker] = useState(null);
  const prevMarker = usePrevious(selectedMarker);
  useEffect(() => {
    if (prevMarker) {
      prevMarker.getElement().style.backgroundImage = getMarkerColor(
        prevMarker.number
      );
    }
    if (selectedMarker) {
      selectedMarker.getElement().style.backgroundImage = `url(${selectMarkerIcon_orange})`;
    }
  }, [selectedMarker]);

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

      const markerObject = new mapboxgl.Marker(el)
        .setLngLat([marker.longitude, marker.latitude])
        .setDraggable(false)
        .addTo(map);

      el.addEventListener("click", () => {
        setSelectedMarker(markerObject);
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

  const getMarkerColor1 = (watingNum) => {
    if (watingNum <= 10) {
      return "green";
    } else if (watingNum <= 20) {
      return "orange";
    } else {
      return "red";
    }
  };
  const getMarkerText = (watingNum) => {
    if (watingNum <= 10) {
      return "여유";
    } else if (watingNum <= 20) {
      return "보통";
    } else {
      return "혼잡";
    }
  };

  return (
    <div>
      <div id="map" style={{ width: "100vw", height: "100vh" }}></div>
      <div>
        <div style={{ position: "relative" }}>
          <button
            style={{
              backgroundColor: "transparent",
              position: "fixed",
              top: "680px",
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
              backgroundColor: "transparent",
              position: "fixed",
              top: "680px",
              right: "5px",
              margin: "0px",
              padding: "0px",
              border: "0px",
            }}
          >
            <img
              src={writeButtonIcon}
              alt="Write Button Icon"
              style={{ width: "72px" }}
            />
          </button>
        </div>
      </div>
    </div>
  );
}

UnSelectedCase.defaultProps = {
  name: "꿈나무소아청소년과의원",
  location: "대전과역시 유성구 전민동 엑스포로 꿈나무소아청소년과의원",
  etc: "#전문의벼원 #주말진료 ",
  watingNum: 30,
  watingTime: "1시간 30분",
  distance: 300,
};

export default UnSelectedCase;
