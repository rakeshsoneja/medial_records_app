import React, { useState, useEffect } from 'react';
import AppointmentController from '../controllers/AppointmentController';
import AppointmentForm from '../components/forms/AppointmentForm';
import AppointmentCard from '../components/cards/AppointmentCard';
import './Records.css';

const Appointments = () => {
  const [appointments, setAppointments] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState({
    doctor_name: '',
    specialty: '',
    hospital: '',
    location: '',
    appointment_date: '',
    notes: '',
  });

  useEffect(() => {
    fetchAppointments();
  }, []);

  const fetchAppointments = async () => {
    setLoading(true);
    const result = await AppointmentController.fetchAll(10, 0, true);
    if (result.success) {
      setAppointments(result.data);
    } else {
      alert(result.error);
    }
    setLoading(false);
  };

  const handleSubmit = async () => {
    const result = await AppointmentController.create(formData);
    if (result.success) {
      setShowForm(false);
      setFormData({
        doctor_name: '',
        specialty: '',
        hospital: '',
        location: '',
        appointment_date: '',
        notes: '',
      });
      fetchAppointments();
    } else {
      alert(result.error);
    }
  };

  if (loading) {
    return <div className="loading">Loading appointments...</div>;
  }

  return (
    <div className="records-page">
      <div className="page-header">
        <h1>Appointments</h1>
        <button className="btn btn-primary" onClick={() => setShowForm(!showForm)}>
          {showForm ? 'Cancel' : 'Add Appointment'}
        </button>
      </div>

      {showForm && (
        <AppointmentForm
          formData={formData}
          onChange={setFormData}
          onSubmit={handleSubmit}
          onCancel={() => setShowForm(false)}
        />
      )}

      <div className="records-list">
        {appointments.length > 0 ? (
          appointments.map((appointment) => (
            <AppointmentCard key={appointment.id} appointment={appointment} />
          ))
        ) : (
          <p>No upcoming appointments. Add your first appointment above.</p>
        )}
      </div>
      <div className="fixed-bottom-bar">
        <button className="btn btn-primary" onClick={() => setShowForm(!showForm)}>
          {showForm ? 'Cancel' : 'Add Appointment'}
        </button>
      </div>
    </div>
  );
};

export default Appointments;
