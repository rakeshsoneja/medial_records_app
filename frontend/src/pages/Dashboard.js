import React, { useState, useEffect } from 'react';
import DashboardController from '../controllers/DashboardController';
import './Dashboard.css';

const Dashboard = () => {
  const [dashboardData, setDashboardData] = useState({
    prescriptions: [],
    appointments: [],
    labReports: [],
    medications: [],
    reminders: [],
  });
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchDashboard();
  }, []);

  const fetchDashboard = async () => {
    setLoading(true);
    const result = await DashboardController.fetchDashboard();
    if (result.success) {
      setDashboardData(result.data);
    } else {
      console.error('Failed to fetch dashboard:', result.error);
    }
    setLoading(false);
  };

  if (loading) {
    return <div className="loading">Loading dashboard...</div>;
  }

  return (
    <div className="dashboard">
      <h1>Dashboard</h1>

      <div className="dashboard-grid">
        <div className="dashboard-card">
          <h2>Upcoming Appointments</h2>
          {dashboardData.appointments.length > 0 ? (
            <ul>
              {dashboardData.appointments.map((apt) => (
                <li key={apt.id}>
                  <strong>{apt.doctorName}</strong> - {apt.specialty}
                  <br />
                  <small>{apt.formattedDate}</small>
                </li>
              ))}
            </ul>
          ) : (
            <p>No upcoming appointments</p>
          )}
        </div>

        <div className="dashboard-card">
          <h2>Active Prescriptions</h2>
          {dashboardData.prescriptions.length > 0 ? (
            <ul>
              {dashboardData.prescriptions.map((pres) => (
                <li key={pres.id}>
                  <strong>{pres.medicineName}</strong> - {pres.dosage}
                  <br />
                  <small>Dr. {pres.prescribingDoctor}</small>
                </li>
              ))}
            </ul>
          ) : (
            <p>No active prescriptions</p>
          )}
        </div>

        <div className="dashboard-card">
          <h2>Recent Lab Reports</h2>
          {dashboardData.labReports.length > 0 ? (
            <ul>
              {dashboardData.labReports.map((report) => (
                <li key={report.id}>
                  <strong>{report.testType}</strong>
                  <br />
                  <small>{report.labName} - {report.formattedDate}</small>
                </li>
              ))}
            </ul>
          ) : (
            <p>No recent lab reports</p>
          )}
        </div>

        <div className="dashboard-card">
          <h2>Active Medications</h2>
          {dashboardData.medications.length > 0 ? (
            <ul>
              {dashboardData.medications.map((med) => (
                <li key={med.id}>
                  <strong>{med.medicineName}</strong> - {med.dosage}
                  <br />
                  <small>Pharmacy: {med.pharmacyName}</small>
                </li>
              ))}
            </ul>
          ) : (
            <p>No active medications</p>
          )}
        </div>

        <div className="dashboard-card">
          <h2>Upcoming Reminders</h2>
          {dashboardData.reminders.length > 0 ? (
            <ul>
              {dashboardData.reminders.map((reminder) => (
                <li key={reminder.id}>
                  <strong>{reminder.title}</strong>
                  <br />
                  <small>{reminder.formattedDate}</small>
                </li>
              ))}
            </ul>
          ) : (
            <p>No upcoming reminders</p>
          )}
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
