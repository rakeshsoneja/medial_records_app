import React, { useState, useEffect } from 'react';
import ReminderController from '../controllers/ReminderController';
import ReminderForm from '../components/forms/ReminderForm';
import ReminderCard from '../components/cards/ReminderCard';
import './Records.css';

const Reminders = () => {
  const [reminders, setReminders] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState({
    title: '',
    description: '',
    reminder_date: '',
    reminder_type: 'checkup',
    is_recurring: false,
    recurrence_interval: '',
  });

  useEffect(() => {
    fetchReminders();
  }, []);

  const fetchReminders = async () => {
    setLoading(true);
    const result = await ReminderController.fetchAll(true);
    if (result.success) {
      setReminders(result.data);
    } else {
      alert(result.error);
    }
    setLoading(false);
  };

  const handleSubmit = async () => {
    const result = await ReminderController.create(formData);
    if (result.success) {
      setShowForm(false);
      setFormData({
        title: '',
        description: '',
        reminder_date: '',
        reminder_type: 'checkup',
        is_recurring: false,
        recurrence_interval: '',
      });
      fetchReminders();
    } else {
      alert(result.error);
    }
  };

  if (loading) {
    return <div className="loading">Loading reminders...</div>;
  }

  return (
    <div className="records-page">
      <div className="page-header">
        <h1>Health Reminders</h1>
        <button className="btn btn-primary" onClick={() => setShowForm(!showForm)}>
          {showForm ? 'Cancel' : 'Add Reminder'}
        </button>
      </div>

      {showForm && (
        <ReminderForm
          formData={formData}
          onChange={setFormData}
          onSubmit={handleSubmit}
          onCancel={() => setShowForm(false)}
        />
      )}

      <div className="records-list">
        {reminders.length > 0 ? (
          reminders.map((reminder) => (
            <ReminderCard key={reminder.id} reminder={reminder} />
          ))
        ) : (
          <p>No reminders found. Add your first reminder above.</p>
        )}
      </div>
    </div>
  );
};

export default Reminders;
