import { useEffect, useState } from "react";
import "./ChatField.css";

function ChatField() {
  const [Message, setMessage] = useState("");
  const [Messages, setMessages] = useState<string[]>([]);

  function handleKeyDown(e: any) {
    if (e.key == "Enter" && Message.trim() != "") {
      setMessages((prev) => [...prev, Message]);
      setMessage("");
      console.log(Messages);
    }
  }

  // auto scroll to the bottom
  useEffect(() => {
    const box = document.getElementById("bust");
    if (box) {
      box.scrollTop = box.scrollHeight;
    }
  }, [Messages]);

  return (
    <div className="bigjas">
      <div className="littlejas" id="bust">
        {Messages.map((msg, i) => (
          <div key={i} className="chat-message">
            {msg}
          </div>
        ))}
      </div>
      <div className="text-wrapper">
        <input
          className="text-field"
          name="myInput"
          value={Message}
          placeholder="TYPE"
          onChange={(e) => setMessage(e.target.value)}
          onKeyDown={(e) => handleKeyDown(e)}
        />
      </div>
    </div>
  );
}

export default ChatField;
