import React from 'react';
import './Card.css';

const LabReportCard = ({ labReport }) => {
  return (
    <div className="card">
      <h3>{labReport.testType}</h3>
      <div className="card-body">
        <p><strong>Lab:</strong> {labReport.labName}</p>
        <p><strong>Test Date:</strong> {labReport.formattedDate}</p>
        {labReport.reportUrl && (
          <p>
            <strong>Report:</strong>{' '}
            <a href={labReport.reportUrl} target="_blank" rel="noopener noreferrer">
              View Report
            </a>
          </p>
        )}
        {labReport.notes && (
          <p><strong>Notes:</strong> {labReport.notes}</p>
        )}
      </div>
    </div>
  );
};

export default LabReportCard;




