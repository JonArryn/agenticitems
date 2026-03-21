import axios from "axios";
import { useState } from "react";

export function App() {
  const [apiData, setApiData] = useState("");

  const handleButtonClick = async () => {
    const { data } = await axios.get("http://localhost:4000/");
    setApiData(data);
  };
  return (
    <main style={{ fontFamily: "system-ui", padding: "1.5rem" }}>
      <h1>itemui</h1>
      <p>
        React UI — API on host port 4000, agents on 4001 (see docker-compose).
      </p>
      <button onClick={handleButtonClick}>Get Server Response</button>
      <div>
        <p>Api Response:</p>
        <p>{!apiData ? "No Api Data Yet" : apiData}</p>{" "}
      </div>
    </main>
  );
}
