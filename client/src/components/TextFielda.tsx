import { useState } from "react";
import "./TextFielda.css";

function TextFielda() {
  const [Message, setMessage] = useState("");

  function handleKeyDown(e: any) {
    if (e.code == "Enter") {
      setMessage("");
    }
  }

  return (
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
  );
}

export default TextFielda;
