import logo from "./logo.svg";
import "./App.css";
import Main from "./Main/Main";
import MapMain from "./Map/MapMain";
import { Route, Routes } from "react-router-dom";
const App = () => {
  return (
    <div>
      <Routes>
        <Route exact path="/" element={<Main />} />
        <Route exact path="/map" element={<MapMain />} />
      </Routes>
    </div>
  );
};

export default App;
