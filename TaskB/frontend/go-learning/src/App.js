import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Navbar from './components/Navbar';
import PersonListPage from './pages/PersonListPage';
import CreatePersonPage from './pages/CreatePersonPage';
import './App.css';

function App() {
  return (
    <Router>
      <Navbar />
      <div className="app-container">
        <Routes>
          <Route path="/" element={<PersonListPage />} />
          <Route path="/create" element={<CreatePersonPage />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
