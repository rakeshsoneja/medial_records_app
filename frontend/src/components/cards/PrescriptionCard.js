import React from 'react';
import './Card.css';

const PrescriptionCard = ({ prescription, onDelete }) => {
  return (
    <div className="card">
      <div className="card-header">
        <h3>{prescription.medicineName}</h3>
        {onDelete && (
          <button
            className="btn btn-danger btn-sm"
            onClick={() => onDelete(prescription.id)}
          >
            Delete
          </button>
        )}
      </div>
      <div className="card-body">
        <p><strong>Dosage:</strong> {prescription.dosage}</p>
        <p><strong>Doctor:</strong> {prescription.prescribingDoctor}</p>
        <p><strong>Specialty:</strong> {prescription.doctorSpecialty}</p>
        <p><strong>Hospital:</strong> {prescription.hospital}</p>
        <p><strong>Date:</strong> {new Date(prescription.prescriptionDate).toLocaleDateString()}</p>
        {prescription.instructions && (
          <p><strong>Instructions:</strong> {prescription.instructions}</p>
        )}
      </div>
    </div>
  );
};

export default PrescriptionCard;




