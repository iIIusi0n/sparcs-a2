import React from "react";
import messageProfile from "../icons/messageprofile.svg";

const MessageListElement = (message) => {
  return (
    <div style={{ display: "flex" }}>
      <img src={messageProfile} alt="Message Profile" />
      <div>{message}</div>
    </div>
  );
};

export default MessageListElement;
