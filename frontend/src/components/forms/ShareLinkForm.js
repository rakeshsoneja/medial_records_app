import React from 'react';
import './Form.css';

const ShareLinkForm = ({ formData, onChange, onSubmit, onCancel }) => {
  const handleChange = (field, value) => {
    onChange({ ...formData, [field]: value });
  };

  return (
    <div className="card">
      <h2>Create Shareable Link</h2>
      <form onSubmit={(e) => { e.preventDefault(); onSubmit(); }}>
        <div className="form-group">
          <label>Record Type *</label>
          <select
            value={formData.record_type || 'prescription'}
            onChange={(e) => handleChange('record_type', e.target.value)}
            required
          >
            <option value="prescription">Prescription</option>
            <option value="appointment">Appointment</option>
            <option value="lab_report">Lab Report</option>
            <option value="bundle">Bundle (Multiple)</option>
          </select>
        </div>
        <div className="form-group">
          <label>Record IDs (comma-separated) *</label>
          <input
            type="text"
            value={formData.record_ids || ''}
            onChange={(e) => handleChange('record_ids', e.target.value)}
            placeholder="e.g., uuid1, uuid2, uuid3"
            required
          />
        </div>
        <div className="form-group">
          <label>Expires In (hours) *</label>
          <input
            type="number"
            value={formData.expires_in_hours || 24}
            onChange={(e) => handleChange('expires_in_hours', parseInt(e.target.value))}
            min="1"
            required
          />
        </div>
        <div className="form-group">
          <label>Max Access Count (0 = unlimited)</label>
          <input
            type="number"
            value={formData.max_access_count || 0}
            onChange={(e) => handleChange('max_access_count', parseInt(e.target.value))}
            min="0"
          />
        </div>
        <div className="form-group">
          <label>Share Method *</label>
          <select
            value={formData.share_method || 'link'}
            onChange={(e) => handleChange('share_method', e.target.value)}
            required
          >
            <option value="link">Link</option>
            <option value="email">Email</option>
            <option value="sms">SMS</option>
          </select>
        </div>
        {formData.share_method === 'email' && (
          <div className="form-group">
            <label>Recipient Email *</label>
            <input
              type="email"
              value={formData.recipient_email || ''}
              onChange={(e) => handleChange('recipient_email', e.target.value)}
              required
            />
          </div>
        )}
        {formData.share_method === 'sms' && (
          <div className="form-group">
            <label>Recipient Phone *</label>
            <input
              type="tel"
              value={formData.recipient_phone || ''}
              onChange={(e) => handleChange('recipient_phone', e.target.value)}
              required
            />
          </div>
        )}
        <div className="form-group">
          <label>
            <input
              type="checkbox"
              checked={formData.allow_download || false}
              onChange={(e) => handleChange('allow_download', e.target.checked)}
            />
            Allow Download
          </label>
        </div>
        <div className="form-actions">
          <button type="submit" className="btn btn-primary">Create Share Link</button>
          <button type="button" className="btn btn-secondary" onClick={onCancel}>Cancel</button>
        </div>
      </form>
    </div>
  );
};

export default ShareLinkForm;




