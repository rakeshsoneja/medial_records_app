import React from 'react';
import './Card.css';

const AppointmentCard = ({ appointment }) => {
  return (
    <div className="card">
      <h3>{appointment.doctorName}</h3>
      <div className="card-body">
        <p><strong>Specialty:</strong> {appointment.specialty}</p>
        <p><strong>Hospital:</strong> {appointment.hospital}</p>
        <p><strong>Location:</strong> {appointment.location}</p>
        <p><strong>Date & Time:</strong> {appointment.formattedDate}</p>
        {appointment.notes && (
          <p><strong>Notes:</strong> {appointment.notes}</p>
        )}
      </div>
    </div>
  );
};

export default AppointmentCard;




