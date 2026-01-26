import React from 'react';
import './Form.css';

const InsuranceForm = ({ formData, onChange, onSubmit, onCancel }) => {
  const handleChange = (field, value) => {
    onChange({ ...formData, [field]: value });
  };

  return (
    <div className="card">
      <h2>Add Health Insurance</h2>
      <form onSubmit={(e) => { e.preventDefault(); onSubmit(); }}>
        <div className="form-group">
          <label>Insurance Provider *</label>
          <input
            type="text"
            value={formData.insurance_provider || ''}
            onChange={(e) => handleChange('insurance_provider', e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label>Policy Number *</label>
          <input
            type="text"
            value={formData.policy_number || ''}
            onChange={(e) => handleChange('policy_number', e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label>Group Number</label>
          <input
            type="text"
            value={formData.group_number || ''}
            onChange={(e) => handleChange('group_number', e.target.value)}
          />
        </div>
        <div className="form-group">
          <label>Member ID</label>
          <input
            type="text"
            value={formData.member_id || ''}
            onChange={(e) => handleChange('member_id', e.target.value)}
          />
        </div>
        <div className="form-group">
          <label>Effective Date</label>
          <input
            type="date"
            value={formData.effective_date || ''}
            onChange={(e) => handleChange('effective_date', e.target.value)}
          />
        </div>
        <div className="form-group">
          <label>Expiration Date</label>
          <input
            type="date"
            value={formData.expiration_date || ''}
            onChange={(e) => handleChange('expiration_date', e.target.value)}
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

export default InsuranceForm;




