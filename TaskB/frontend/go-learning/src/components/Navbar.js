import React from 'react';
import { Link } from 'react-router-dom';
import './Navbar.css';

function Navbar() {
  return (
    <header className="navbar">
      <div className="navbar-container">
        <nav className="nav">
          <Link to="/" className="nav-link">Persons</Link>
          <Link to="/create" className="nav-link">Create Person</Link>
          {/* You can add more links here like "Reports", "Orders", etc. */}
        </nav>
      </div>
    </header>
  );
}

export default Navbar;
