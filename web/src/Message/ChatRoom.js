import React, { useState } from "react";
import { useEffect } from "react";
import axios from "axios";
import ChatrooomHeader from "../components/chatroom/ChatroomHeader.jsx";
import NavigationBarChat from "../components/navigation/NavigationBarChat.jsx";
import sendIcon from "../components/icons/send.svg";
import addIcon from "../components/icons/add.svg";

const ChatRoom = (props) => {
  const [messages, setMessages] = useState([]);
  const [newMessage, setNewMessage] = useState("");

  useEffect(() => {
    fetchMessages();
  }, []);

  const fetchMessages = async () => {
    try {
      const response = await axios.get("/api/messages");
      setMessages(response.data);
    } catch (error) {
      console.error("Error fetching messages:", error);
    }
  };

  const sendMessage = async () => {
    try {
      await axios.post("/api/messages", { text: newMessage });
      setNewMessage("");
      fetchMessages();
    } catch (error) {
      console.error("Error sending message:", error);
    }
  };
  return (
    <div>
      <ChatrooomHeader></ChatrooomHeader>
      <div style={{ position: "relative", top: "50px" }}>
        <h1>Message Board</h1>
        <div
          style={{
            width: "100vw",
            height: "6vh",
            backgroundColor: "#D1D2D3",
            position: "absolute",
            top: "1450%",
          }}
        >
          <img
            src={addIcon}
            alt="Add Icon"
            style={{ position: "absolute", top: "30%", left: "5%" }}
          />
          <input
            type="text"
            value={newMessage}
            onChange={(e) => setNewMessage(e.target.value)}
            style={{
              borderRadius: "10px",
              border: "none",
              height: "3vh",
              width: "70vw",
              position: "absolute",
              left: "15%",
              top: "30%",
            }}
          />
          <button
            onClick={sendMessage}
            style={{
              background: "#D1D2D3",
              border: "none",
              position: "absolute",
              top: "33%",
              left: "88%",
            }}
          >
            <img src={sendIcon} alt="Send Icon" />
          </button>
        </div>
        <div>
          {messages.map((message, index) => (
            <div key={index}>{message.text}</div>
          ))}
        </div>
      </div>
      <NavigationBarChat></NavigationBarChat>
    </div>
  );
};
export default ChatRoom;
