import React from "react";
import "./App.css";
import { container } from "semantic-ui-react";
import ToDo from "./To-Do";

function App() {
  return (
    <div>
      <container>
        <ToDo />
      </container>
    </div>
  );
}

export default App;
