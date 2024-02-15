import React, { useState } from "react";
import ChatrooomHeader from "../components/chatroom/ChatroomHeader.jsx";
import NavigationBarChat from "../components/navigation/NavigationBarChat.jsx";

const ChatRoom = (props) => {
  return (
    <div>
      <ChatrooomHeader></ChatrooomHeader>
      <NavigationBarChat></NavigationBarChat>
    </div>
  );
};
export default ChatRoom;
