import React from 'react';
import Header from "./components/Header";
import Form from "./components/Form";

function App() {
  return (
    <div className="context">
      <Header />
      <main className="mt-14 max-w-xl mx-auto divide-y md:max-w-4x1">
        <Form />
      </main>
    </div>
  );
}

export default App;
