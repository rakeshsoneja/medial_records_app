import React from 'react';
import './Form.css';

const PrescriptionForm = ({ formData, onChange, onSubmit, onCancel }) => {
  const handleChange = (field, value) => {
    onChange({ ...formData, [field]: value });
  };

  return (
    <div className="card">
      <h2>Add New Prescription</h2>
      <form onSubmit={(e) => { e.preventDefault(); onSubmit(); }}>
        <div className="form-group">
          <label>Medicine Name *</label>
          <input
            type="text"
            value={formData.medicine_name || ''}
            onChange={(e) => handleChange('medicine_name', e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label>Dosage</label>
          <input
            type="text"
            value={formData.dosage || ''}
            onChange={(e) => handleChange('dosage', e.target.value)}
          />
        </div>
        <div className="form-group">
          <label>Instructions</label>
          <textarea
            value={formData.instructions || ''}
            onChange={(e) => handleChange('instructions', e.target.value)}
            rows="3"
          />
        </div>
        <div className="form-group">
          <label>Prescribing Doctor</label>
          <input
            type="text"
            value={formData.prescribing_doctor || ''}
            onChange={(e) => handleChange('prescribing_doctor', e.target.value)}
          />
        </div>
        <div className="form-group">
          <label>Doctor Specialty</label>
          <input
            type="text"
            value={formData.doctor_specialty || ''}
            onChange={(e) => handleChange('doctor_specialty', e.target.value)}
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
          <label>Prescription Date *</label>
          <input
            type="date"
            value={formData.prescription_date || new Date().toISOString().split('T')[0]}
            onChange={(e) => handleChange('prescription_date', e.target.value)}
            required
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

export default PrescriptionForm;




