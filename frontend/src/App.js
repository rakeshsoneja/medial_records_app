import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { AuthProvider } from './contexts/AuthContext';
import PrivateRoute from './components/PrivateRoute';
import Login from './pages/Login';
import Register from './pages/Register';
import Dashboard from './pages/Dashboard';
import Prescriptions from './pages/Prescriptions';
import Appointments from './pages/Appointments';
import LabReports from './pages/LabReports';
import Medications from './pages/Medications';
import Reminders from './pages/Reminders';
import Insurance from './pages/Insurance';
import Sharing from './pages/Sharing';
import Layout from './components/Layout';
import './App.css';

function App() {
  return (
    <AuthProvider>
      <Router>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route
            path="/"
            element={
              <PrivateRoute>
                <Layout>
                  <Dashboard />
                </Layout>
              </PrivateRoute>
            }
          />
          <Route
            path="/prescriptions"
            element={
              <PrivateRoute>
                <Layout>
                  <Prescriptions />
                </Layout>
              </PrivateRoute>
            }
          />
          <Route
            path="/appointments"
            element={
              <PrivateRoute>
                <Layout>
                  <Appointments />
                </Layout>
              </PrivateRoute>
            }
          />
          <Route
            path="/lab-reports"
            element={
              <PrivateRoute>
                <Layout>
                  <LabReports />
                </Layout>
              </PrivateRoute>
            }
          />
          <Route
            path="/medications"
            element={
              <PrivateRoute>
                <Layout>
                  <Medications />
                </Layout>
              </PrivateRoute>
            }
          />
          <Route
            path="/reminders"
            element={
              <PrivateRoute>
                <Layout>
                  <Reminders />
                </Layout>
              </PrivateRoute>
            }
          />
          <Route
            path="/insurance"
            element={
              <PrivateRoute>
                <Layout>
                  <Insurance />
                </Layout>
              </PrivateRoute>
            }
          />
          <Route
            path="/sharing"
            element={
              <PrivateRoute>
                <Layout>
                  <Sharing />
                </Layout>
              </PrivateRoute>
            }
          />
        </Routes>
      </Router>
    </AuthProvider>
  );
}

export default App;




