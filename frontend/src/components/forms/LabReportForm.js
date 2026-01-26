import React from 'react';
import './Form.css';

const LabReportForm = ({ formData, onChange, onSubmit, onCancel }) => {
  const handleChange = (field, value) => {
    onChange({ ...formData, [field]: value });
  };

  return (
    <div className="card">
      <h2>Add New Lab Report</h2>
      <form onSubmit={(e) => { e.preventDefault(); onSubmit(); }}>
        <div className="form-group">
          <label>Test Type *</label>
          <input
            type="text"
            value={formData.test_type || ''}
            onChange={(e) => handleChange('test_type', e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label>Lab Name</label>
          <input
            type="text"
            value={formData.lab_name || ''}
            onChange={(e) => handleChange('lab_name', e.target.value)}
          />
        </div>
        <div className="form-group">
          <label>Test Date *</label>
          <input
            type="date"
            value={formData.test_date || new Date().toISOString().split('T')[0]}
            onChange={(e) => handleChange('test_date', e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label>Report URL (after upload)</label>
          <input
            type="text"
            value={formData.report_url || ''}
            onChange={(e) => handleChange('report_url', e.target.value)}
            placeholder="Will be populated after file upload"
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

export default LabReportForm;




