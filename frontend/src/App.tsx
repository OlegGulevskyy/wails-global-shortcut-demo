import { useEffect, useState } from "react";
import logo from "./assets/images/logo-universal.png";
import "./App.css";
import { Greet } from "../wailsjs/go/main/App";
import { EventsOn } from "../wailsjs/runtime/runtime";

function App() {
  const [resultText, setResultText] = useState(
    "Please enter your name below 👇"
  );
  const [name, setName] = useState("");
  const updateName = (e: any) => setName(e.target.value);
  const updateResultText = (result: string) => setResultText(result);
  const [currentTime, setCurrentTime] = useState<string | null>(null);

  function globalHotkeyEventHandler(time: string) {
    setCurrentTime(time);
  }

  useEffect(() => {
    EventsOn("Backend:GlobalHotkeyEvent", globalHotkeyEventHandler);
  }, []);

  function greet() {
    Greet(name).then(updateResultText);
  }

  const time =
    currentTime ||
    "Use keyboard shortcut Ctrl + Shift + S to find out current time";

  return (
    <div id="App">
      <img src={logo} id="logo" alt="logo" />
      <div id="result" className="result">
        {resultText}
      </div>
      <div id="input" className="input-box">
        <input
          id="name"
          className="input"
          onChange={updateName}
          autoComplete="off"
          name="input"
          type="text"
        />
        <button className="btn" onClick={greet}>
          Greet
        </button>
        <div>
          Current time from backend:
          <p>{time}</p>
        </div>
        <p className="disclaimer">
          This keyboard shortcut works anywhere on your PC, as long as the app
          process is running, even if the window is not focused.
        </p>
      </div>
    </div>
  );
}

export default App;
