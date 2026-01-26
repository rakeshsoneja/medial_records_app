import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { FaFileMedical } from 'react-icons/fa';
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
          <button className="nav-toggle" onClick={() => setSidebarOpen(!sidebarOpen)} aria-label="Menu">
            ☰
          </button>
          <h1 className="nav-logo">
            <FaFileMedical className="nav-logo-icon" />
            <span>Health Portal</span>
          </h1>
          <div className="nav-user">
            <span className="nav-user-name">{user?.first_name} {user?.last_name}</span>
            <button onClick={handleLogout} className="nav-logout-btn nav-logout-btn-desktop" aria-label="Logout">
              <span className="logout-icon">↪</span>
            </button>
          </div>
        </div>
      </nav>

      <div className="layout-content">
        <aside className={`sidebar ${sidebarOpen ? 'open' : ''}`}>
          <nav className="sidebar-nav">
            <Link to="/" onClick={() => setSidebarOpen(false)}>
              <span className="sidebar-icon">⚡</span>
              <span>Dashboard</span>
            </Link>
            <Link to="/prescriptions" onClick={() => setSidebarOpen(false)}>
              <span className="sidebar-icon">⚕</span>
              <span>Prescriptions</span>
            </Link>
            <Link to="/appointments" onClick={() => setSidebarOpen(false)}>
              <span className="sidebar-icon">📅</span>
              <span>Appointments</span>
            </Link>
            <Link to="/lab-reports" onClick={() => setSidebarOpen(false)}>
              <span className="sidebar-icon">⚗</span>
              <span>Lab Reports</span>
            </Link>
            <Link to="/medications" onClick={() => setSidebarOpen(false)}>
              <span className="sidebar-icon">💊</span>
              <span>Medications</span>
            </Link>
            <Link to="/reminders" onClick={() => setSidebarOpen(false)}>
              <span className="sidebar-icon">⏰</span>
              <span>Reminders</span>
            </Link>
            <Link to="/insurance" onClick={() => setSidebarOpen(false)}>
              <span className="sidebar-icon">🛡</span>
              <span>Insurance</span>
            </Link>
            <Link to="/sharing" onClick={() => setSidebarOpen(false)}>
              <span className="sidebar-icon">🔗</span>
              <span>Sharing</span>
            </Link>
          </nav>
          <div className="sidebar-footer">
            <div className="sidebar-separator"></div>
            <div className="sidebar-user-info">
              <span className="sidebar-user-icon">👤</span>
              <span className="sidebar-user-name-text">{user?.first_name} {user?.last_name}</span>
            </div>
            <button className="sidebar-logout-item" onClick={handleLogout}>
              <span className="sidebar-logout-icon">↪</span>
              <span>Logout</span>
            </button>
          </div>
          <div className="sidebar-footer-desktop">
            <div className="sidebar-separator-desktop"></div>
            <div className="sidebar-user-info-desktop">
              <span className="sidebar-user-icon-desktop">👤</span>
              <span className="sidebar-user-name-text-desktop">{user?.first_name} {user?.last_name}</span>
            </div>
            <button className="sidebar-logout-item-desktop" onClick={handleLogout}>
              <span className="sidebar-logout-icon-desktop">↪</span>
              <span>Logout</span>
            </button>
          </div>
        </aside>

        <main className="main-content">
          {children}
        </main>
      </div>
    </div>
  );
};

export default Layout;




