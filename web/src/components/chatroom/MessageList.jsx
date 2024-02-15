import React from "react";
import MessageListElement from "./MessageListElement";

function MessageList(messages) {
  return (
    <div>
      <h2>Messages:</h2>
      <ul>
        {messages.map((message, index) => (
          <MessageListElement message={message}></MessageListElement>
        ))}
      </ul>
    </div>
  );
}

export default MessageList;
