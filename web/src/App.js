import logo from "./logo.svg";
import "./App.css";
import Main from "./Main/Main";
import MapMain from "./Map/MapMain";
import { Route, Routes } from "react-router-dom";
import MessageMain from "./Message/MessageMain";
import ChatRoom from "./Message/ChatRoom";
import AddNumber from "./AddNumber/AddNumber";
import Location from "./components/Location";
const App = () => {
  return (
    <div>
      <Routes>
        <Route exact path="/" element={<Main />} />
        <Route exact path="/map" element={<MapMain />} />
        <Route exact path="/message" element={<MessageMain />} />
        <Route exact path="/chatroom" element={<ChatRoom />} />
        <Route exact path="/AddNumber" element={<AddNumber />} />
      </Routes>
    </div>
  );
};

export default App;
