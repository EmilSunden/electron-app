import React, { useEffect, useState } from "react";

import ReadAllApps from "./components/readAllApps/ReadAllApps";
import FieldInput from "./components/fieldInput/FieldInput";

function App() {
  return (
    <div className="App">
      <FieldInput />
      <ReadAllApps/>
    </div>
  );
}

export default App;
