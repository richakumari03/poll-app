import React from 'react';
//import './App.css';
import AppRouter from './app.router';
import Navbar from './pages/Navbar';

function App() {
  return (
    <div className="container">
      <Navbar />
      <AppRouter />
    </div>
  );
}

export default App;