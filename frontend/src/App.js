import React from "react";

import ReadAllApps from "./components/readAllApps/ReadAllApps";
import FieldInput from "./components/fieldInput/FieldInput";
import Container from "./components/container/Container";

function App() {
  return (
    <Container>
      <FieldInput />
      <ReadAllApps />
    </Container>
  );
}

export default App;
