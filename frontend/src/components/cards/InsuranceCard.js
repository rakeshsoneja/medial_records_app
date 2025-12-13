import React from 'react';
import './Card.css';

const InsuranceCard = ({ insurance }) => {
  return (
    <div className="card">
      <h3>{insurance.insuranceProvider}</h3>
      <div className="card-body">
        <p><strong>Policy Number:</strong> {insurance.policyNumber}</p>
        {insurance.groupNumber && (
          <p><strong>Group Number:</strong> {insurance.groupNumber}</p>
        )}
        {insurance.memberId && (
          <p><strong>Member ID:</strong> {insurance.memberId}</p>
        )}
        {insurance.effectiveDate && (
          <p>
            <strong>Effective Date:</strong>{' '}
            {new Date(insurance.effectiveDate).toLocaleDateString()}
          </p>
        )}
        {insurance.expirationDate && (
          <p>
            <strong>Expiration Date:</strong>{' '}
            {new Date(insurance.expirationDate).toLocaleDateString()}
          </p>
        )}
        {insurance.notes && (
          <p><strong>Notes:</strong> {insurance.notes}</p>
        )}
      </div>
    </div>
  );
};

export default InsuranceCard;

