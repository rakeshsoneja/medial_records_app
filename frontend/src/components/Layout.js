import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import './Layout.css';

const Layout = ({ children }) => {
  const { user, logout } = useAuth();
  const navigate = useNavigate();
  const [sidebarOpen, setSidebarOpen] = useState(false);

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  return (
    <div className="layout">
      <nav className="navbar">
        <div className="nav-container">
          <h1 className="nav-logo">Medical Records</h1>
          <button className="nav-toggle" onClick={() => setSidebarOpen(!sidebarOpen)}>
            ☰
          </button>
          <div className="nav-user">
            <span>{user?.first_name} {user?.last_name}</span>
            <button onClick={handleLogout} className="btn btn-secondary">Logout</button>
          </div>
        </div>
      </nav>

      <div className="layout-content">
        <aside className={`sidebar ${sidebarOpen ? 'open' : ''}`}>
          <nav className="sidebar-nav">
            <Link to="/" onClick={() => setSidebarOpen(false)}>Dashboard</Link>
            <Link to="/prescriptions" onClick={() => setSidebarOpen(false)}>Prescriptions</Link>
            <Link to="/appointments" onClick={() => setSidebarOpen(false)}>Appointments</Link>
            <Link to="/lab-reports" onClick={() => setSidebarOpen(false)}>Lab Reports</Link>
            <Link to="/medications" onClick={() => setSidebarOpen(false)}>Medications</Link>
            <Link to="/reminders" onClick={() => setSidebarOpen(false)}>Reminders</Link>
            <Link to="/insurance" onClick={() => setSidebarOpen(false)}>Insurance</Link>
            <Link to="/sharing" onClick={() => setSidebarOpen(false)}>Sharing</Link>
          </nav>
        </aside>

        <main className="main-content">
          {children}
        </main>
      </div>
    </div>
  );
};

export default Layout;

