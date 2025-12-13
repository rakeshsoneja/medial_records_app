import React from 'react';
import './Card.css';

const ReminderCard = ({ reminder }) => {
  return (
    <div className="card">
      <h3>{reminder.title}</h3>
      <div className="card-body">
        <p><strong>Type:</strong> {reminder.reminderType}</p>
        <p><strong>Date:</strong> {reminder.formattedDate}</p>
        {reminder.description && (
          <p><strong>Description:</strong> {reminder.description}</p>
        )}
        {reminder.isRecurring && (
          <p><strong>Recurring:</strong> {reminder.recurrenceInterval}</p>
        )}
      </div>
    </div>
  );
};

export default ReminderCard;

