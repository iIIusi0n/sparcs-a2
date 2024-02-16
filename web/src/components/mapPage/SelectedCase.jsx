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
import registerIcon from "../icons/register.svg";
import { useNavigate } from "react-router-dom";

function SelectedCase(props) {
  const [register, setRegister] = useState(false);
  const { name, location, etc, watingNum, watingTime, distance, image } = props;
  const [markers, setMarkers] = useState([
    { longitude: 127.3845475, latitude: 36.3505119, number: 22 },
    { longitude: 127.3822485, latitude: 36.3482102, number: 2 },
    { longitude: 127.3813485, latitude: 36.3493102, number: 13 },
    { longitude: 127.3843485, latitude: 36.3433102, number: 13 },
  ]);

  const toggleRegister = () => {
    localStorage.setItem("hospital", props.name);
    setRegister((prevState) => !prevState);
  };

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
      <div style={{ position: "relative" }}>
        <div id="map" style={{ width: "100vw", height: "100vh" }}></div>
        <div>
          <div>
            <button
              style={{
                backgroundColor: "transparent",
                position: "fixed",
                top: "55%",
                left: "2%",
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
                top: "54%",
                right: "2%",
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
          {register && (
            <div
              style={{
                backgroundColor: "white",
                width: "100vw",
                height: "30vh",
                position: "absolute",
                bottom: "5%",
              }}
            >
              <div style={{ position: "absolute", left: "40%", top: "15%" }}>
                <img src={registerIcon} alt="Current Button Icon" />
              </div>
              <p
                style={{
                  position: "absolute",
                  left: "28%",
                  top: "37%",
                  size: "20px",
                  fontWeight: "bold",
                  fontFamily: "Pretendard-Bold",
                }}
              >
                내 병원으로 등록되었어요!
              </p>
              <button
                style={{
                  backgroundColor: "#FF772A",
                  borderRadius: "10px",
                  width: "80vw",
                  left: "10%",
                  height: "5vh",
                  position: "absolute",
                  top: "60%",
                  border: "none",
                }}
                onClick={toggleRegister}
              >
                <p
                  style={{
                    color: "white",
                    fontWeight: "bold",
                    fontFamily: "Pretendard-Medium",
                  }}
                >
                  확인
                </p>
              </button>
            </div>
          )}
          {!register && (
            <div
              style={{
                backgroundColor: "white",
                width: "100vw",
                height: "20vh",
                paddingBottom: "30%",
                position: "absolute",
                bottom: "1%",
              }}
            >
              <div
                style={{
                  display: "flex",
                  position: "relative",
                  padding: "0px",
                  left: "20px",
                }}
              >
                <img src={hospitalIcon} alt="Hospital Icon" />
                <p
                  style={{
                    fontSize: "20px",
                    fontFamily: "Pretendard-Bold",
                  }}
                >
                  {props.name}
                </p>
                <div
                  style={{
                    backgroundColor: getMarkerColor1(props.watingNum),
                    width: "7vw",
                    height: "4vw",
                    borderRadius: "5px",
                    top: "25px",
                    left: "260px",
                    position: "absolute",
                    fontSize: "12px",
                    color: "white",
                    fontFamily: "pretendatd-Medium",
                  }}
                >
                  {getMarkerText(props.watingNum)}
                </div>
              </div>
              <div
                style={{
                  width: "90px",
                  height: "90px",
                  borderRadius: "10px",
                  backgroundColor: "#A3A5A8",
                  overflow: "hidden",
                  position: "absolute",
                  left: "20px",
                  top: "30%",
                  margin: "0px",
                  padding: "0px",
                }}
              >
                <img src={props.image} alt="이미지" />
              </div>
              <div style={{ display: "flex" }}>
                <p
                  style={{
                    color: "#FF772A",
                    top: "13%",
                    left: "120px",
                    position: "absolute",
                    fontSize: "15px",
                    fontFamily: "pretendard-Bold",
                  }}
                >
                  주소
                </p>
                <p
                  style={{
                    left: "120px",
                    position: "absolute",
                    top: "20%",
                    fontSize: "14px",
                    fontFamily: "pretendard-Medium",
                  }}
                >
                  {props.location}
                </p>
              </div>
              <div style={{ display: "flex" }}>
                <p
                  style={{
                    color: "#FF772A",
                    top: "35%",
                    left: "120px",
                    position: "absolute",
                    fontSize: "15px",
                    fontFamily: "pretendard-Bold",
                  }}
                >
                  기타
                </p>
                <p
                  style={{
                    top: "43%",
                    left: "120px",
                    position: "absolute",
                    fontSize: "14px",
                    fontFamily: "pretendard-Medium",
                  }}
                >
                  {props.etc}
                </p>
              </div>
              <div
                style={{
                  display: "flex",
                  backgroundColor: "#FFF5F0",
                  width: "55vw",
                  height: "6vh",
                  position: "absolute",
                  top: "55%",
                  left: "28%",
                  borderRadius: "15px",
                }}
              >
                <div>
                  <p
                    style={{
                      margin: "0px",
                      color: "#FF772A",
                      marginLeft: "15px",
                      fontSize: "10px",
                    }}
                  >
                    대기번호{" "}
                  </p>
                  <p
                    style={{
                      margin: "0px",
                      position: "absolute",
                      left: "5%",
                    }}
                  >
                    {props.watingNum}번
                  </p>
                </div>
                <div>
                  <p
                    style={{
                      margin: "0px",
                      color: "#FF772A",
                      marginLeft: "23px",
                      fontSize: "10px",
                    }}
                  >
                    대기시간
                  </p>
                  <p
                    style={{
                      margin: "0px",
                      position: "absolute",
                      left: "27%",
                    }}
                  >
                    {props.watingTime}
                  </p>
                </div>
                <div>
                  <p
                    style={{
                      margin: "0px",
                      color: "#FF772A",
                      marginLeft: "45px",
                      fontSize: "10px",
                    }}
                  >
                    {" "}
                    거리{" "}
                  </p>
                  <p
                    style={{
                      margin: "0px",
                      position: "absolute",
                      left: "73%",
                    }}
                  >
                    {props.distance}m
                  </p>
                </div>
              </div>
              <button
                style={{
                  backgroundColor: "#FF772A",
                  borderRadius: "10px",
                  width: "80vw",
                  left: "10%",
                  height: "5vh",
                  position: "absolute",
                  top: "80%",
                  border: "none",
                }}
                onClick={toggleRegister}
              >
                <p style={{ color: "white", fontWeight: "bold" }}>
                  내 병원으로 등록하기
                </p>
              </button>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

SelectedCase.defaultProps = {
  name: "꿈나무소아청소년과의원",
  location: "대전과역시 유성구 전민동 엑스포로 꿈나무소아청소년과의원",
  etc: "#전문의벼원 #주말진료 ",
  watingNum: 30,
  watingTime: "1시간 30분",
  distance: 300,
};

export default SelectedCase;
