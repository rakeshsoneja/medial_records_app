import React from 'react';
import './Card.css';

const MedicationCard = ({ medication }) => {
  return (
    <div className="card">
      <h3>{medication.medicineName}</h3>
      <div className="card-body">
        <p><strong>Dosage:</strong> {medication.dosage}</p>
        <p><strong>Frequency:</strong> {medication.frequency}</p>
        <p><strong>Pharmacy:</strong> {medication.pharmacyName}</p>
        {medication.pharmacyPhone && (
          <p><strong>Pharmacy Phone:</strong> {medication.pharmacyPhone}</p>
        )}
        {medication.nextRefillDate && (
          <p>
            <strong>Next Refill:</strong>{' '}
            {new Date(medication.nextRefillDate).toLocaleDateString()}
          </p>
        )}
      </div>
    </div>
  );
};

export default MedicationCard;

