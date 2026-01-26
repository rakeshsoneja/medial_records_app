import React from 'react';
import './Form.css';

const AppointmentForm = ({ formData, onChange, onSubmit, onCancel }) => {
  const handleChange = (field, value) => {
    onChange({ ...formData, [field]: value });
  };

  return (
    <div className="card">
      <h2>Add New Appointment</h2>
      <form onSubmit={(e) => { e.preventDefault(); onSubmit(); }}>
        <div className="form-group">
          <label>Doctor Name *</label>
          <input
            type="text"
            value={formData.doctor_name || ''}
            onChange={(e) => handleChange('doctor_name', e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label>Specialty</label>
          <input
            type="text"
            value={formData.specialty || ''}
            onChange={(e) => handleChange('specialty', e.target.value)}
          />
        </div>
        <div className="form-group">
          <label>Hospital</label>
          <input
            type="text"
            value={formData.hospital || ''}
            onChange={(e) => handleChange('hospital', e.target.value)}
          />
        </div>
        <div className="form-group">
          <label>Location</label>
          <input
            type="text"
            value={formData.location || ''}
            onChange={(e) => handleChange('location', e.target.value)}
          />
        </div>
        <div className="form-group">
          <label>Appointment Date & Time *</label>
          <input
            type="datetime-local"
            value={formData.appointment_date || ''}
            onChange={(e) => handleChange('appointment_date', e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label>Notes</label>
          <textarea
            value={formData.notes || ''}
            onChange={(e) => handleChange('notes', e.target.value)}
            rows="3"
          />
        </div>
        <div className="form-actions">
          <button type="submit" className="btn btn-primary">Save</button>
          <button type="button" className="btn btn-secondary" onClick={onCancel}>Cancel</button>
        </div>
      </form>
    </div>
  );
};

export default AppointmentForm;




