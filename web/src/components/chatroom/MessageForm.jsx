import React, { useState } from "react";

function MessageForm({ onSendMessage }) {
  const [inputMessage, setInputMessage] = useState("");

  const handleChange = (event) => {
    setInputMessage(event.target.value);
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    onSendMessage(inputMessage);
    setInputMessage("");
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        value={inputMessage}
        onChange={handleChange}
        placeholder="Type your message..."
      />
      <button type="submit">Send</button>
    </form>
  );
}

export default MessageForm;
