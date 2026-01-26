import React from 'react';
import './Form.css';

const MedicationForm = ({ formData, onChange, onSubmit, onCancel }) => {
  const handleChange = (field, value) => {
    onChange({ ...formData, [field]: value });
  };

  return (
    <div className="card">
      <h2>Add New Medication</h2>
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
          <label>Frequency</label>
          <input
            type="text"
            value={formData.frequency || ''}
            onChange={(e) => handleChange('frequency', e.target.value)}
            placeholder="e.g., Daily, Twice daily"
          />
        </div>
        <div className="form-group">
          <label>Pharmacy Name</label>
          <input
            type="text"
            value={formData.pharmacy_name || ''}
            onChange={(e) => handleChange('pharmacy_name', e.target.value)}
          />
        </div>
        <div className="form-group">
          <label>Pharmacy Phone</label>
          <input
            type="tel"
            value={formData.pharmacy_phone || ''}
            onChange={(e) => handleChange('pharmacy_phone', e.target.value)}
          />
        </div>
        <div className="form-group">
          <label>Pharmacy Address</label>
          <input
            type="text"
            value={formData.pharmacy_address || ''}
            onChange={(e) => handleChange('pharmacy_address', e.target.value)}
          />
        </div>
        <div className="form-group">
          <label>Refill Reminder (days before)</label>
          <input
            type="number"
            value={formData.refill_reminder_days || 7}
            onChange={(e) => handleChange('refill_reminder_days', parseInt(e.target.value))}
            min="1"
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

export default MedicationForm;




