import React from 'react';
import './Card.css';

const SharedRecordCard = ({ sharedRecord, onRevoke }) => {
  return (
    <div className="card">
      <div className="card-header">
        <h3>Share Link: {sharedRecord.shareToken?.substring(0, 20)}...</h3>
        {sharedRecord.isActive && onRevoke && (
          <button
            className="btn btn-danger btn-sm"
            onClick={() => onRevoke(sharedRecord.id)}
          >
            Revoke
          </button>
        )}
      </div>
      <div className="card-body">
        <p><strong>Record Type:</strong> {sharedRecord.recordType}</p>
        <p><strong>Expires:</strong> {new Date(sharedRecord.expiresAt).toLocaleString()}</p>
        <p>
          <strong>Access Count:</strong> {sharedRecord.currentAccessCount} /{' '}
          {sharedRecord.maxAccessCount || '∞'}
        </p>
        <p><strong>Status:</strong> {sharedRecord.isActive ? 'Active' : 'Revoked'}</p>
        <p><strong>Share URL:</strong> {sharedRecord.shareUrl}</p>
      </div>
    </div>
  );
};

export default SharedRecordCard;




