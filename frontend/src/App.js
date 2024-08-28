import React, {useEffect, useState} from 'react';

function App() {
  const [message, setMessage] = useState('');

  useEffect(() => {
    const fetchApiUrl = async () => {
      const res = await fetch("http://localhost:8080/api/hello");
      const data = await res.json();
      setMessage(data.message);
    
    }
    fetchApiUrl();
  } , []);

  return (
    <div className="App">
      <h1>{message}</h1>
    </div>
  );
}

export default App;
