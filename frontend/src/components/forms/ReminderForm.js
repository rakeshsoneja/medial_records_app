import React from 'react';
import './Form.css';

const ReminderForm = ({ formData, onChange, onSubmit, onCancel }) => {
  const handleChange = (field, value) => {
    onChange({ ...formData, [field]: value });
  };

  return (
    <div className="card">
      <h2>Add New Reminder</h2>
      <form onSubmit={(e) => { e.preventDefault(); onSubmit(); }}>
        <div className="form-group">
          <label>Title *</label>
          <input
            type="text"
            value={formData.title || ''}
            onChange={(e) => handleChange('title', e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label>Description</label>
          <textarea
            value={formData.description || ''}
            onChange={(e) => handleChange('description', e.target.value)}
            rows="3"
          />
        </div>
        <div className="form-group">
          <label>Reminder Date & Time *</label>
          <input
            type="datetime-local"
            value={formData.reminder_date || ''}
            onChange={(e) => handleChange('reminder_date', e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label>Reminder Type</label>
          <select
            value={formData.reminder_type || 'checkup'}
            onChange={(e) => handleChange('reminder_type', e.target.value)}
          >
            <option value="checkup">Check-up</option>
            <option value="vaccination">Vaccination</option>
            <option value="test">Test</option>
            <option value="other">Other</option>
          </select>
        </div>
        <div className="form-group">
          <label>
            <input
              type="checkbox"
              checked={formData.is_recurring || false}
              onChange={(e) => handleChange('is_recurring', e.target.checked)}
            />
            Recurring Reminder
          </label>
        </div>
        {formData.is_recurring && (
          <div className="form-group">
            <label>Recurrence Interval</label>
            <select
              value={formData.recurrence_interval || ''}
              onChange={(e) => handleChange('recurrence_interval', e.target.value)}
            >
              <option value="monthly">Monthly</option>
              <option value="quarterly">Quarterly</option>
              <option value="yearly">Yearly</option>
            </select>
          </div>
        )}
        <div className="form-actions">
          <button type="submit" className="btn btn-primary">Save</button>
          <button type="button" className="btn btn-secondary" onClick={onCancel}>Cancel</button>
        </div>
      </form>
    </div>
  );
};

export default ReminderForm;




