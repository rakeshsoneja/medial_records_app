import React, { useState, useEffect } from 'react';
import SharingController from '../controllers/SharingController';
import ShareLinkForm from '../components/forms/ShareLinkForm';
import SharedRecordCard from '../components/cards/SharedRecordCard';
import './Records.css';

const Sharing = () => {
  const [sharedRecords, setSharedRecords] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState({
    record_type: 'prescription',
    record_ids: '',
    expires_in_hours: 24,
    max_access_count: 0,
    allow_download: false,
    recipient_email: '',
    recipient_phone: '',
    share_method: 'link',
  });

  useEffect(() => {
    fetchSharedRecords();
  }, []);

  const fetchSharedRecords = async () => {
    setLoading(true);
    const result = await SharingController.fetchMyShares();
    if (result.success) {
      setSharedRecords(result.data);
    } else {
      alert(result.error);
    }
    setLoading(false);
  };

  const handleSubmit = async () => {
    const result = await SharingController.createShareLink(formData);
    if (result.success) {
      setShowForm(false);
      setFormData({
        record_type: 'prescription',
        record_ids: '',
        expires_in_hours: 24,
        max_access_count: 0,
        allow_download: false,
        recipient_email: '',
        recipient_phone: '',
        share_method: 'link',
      });
      fetchSharedRecords();
    } else {
      alert(result.error);
    }
  };

  const handleRevoke = async (id) => {
    if (window.confirm('Are you sure you want to revoke this share link?')) {
      const result = await SharingController.revokeShareLink(id);
      if (result.success) {
        fetchSharedRecords();
      } else {
        alert(result.error);
      }
    }
  };

  if (loading) {
    return <div className="loading">Loading shared records...</div>;
  }

  return (
    <div className="records-page">
      <div className="page-header">
        <h1>Share Records</h1>
        <button className="btn btn-primary" onClick={() => setShowForm(!showForm)}>
          {showForm ? 'Cancel' : 'Create Share Link'}
        </button>
      </div>

      {showForm && (
        <ShareLinkForm
          formData={formData}
          onChange={setFormData}
          onSubmit={handleSubmit}
          onCancel={() => setShowForm(false)}
        />
      )}

      <div className="records-list">
        {sharedRecords.length > 0 ? (
          sharedRecords.map((shared) => (
            <SharedRecordCard
              key={shared.id}
              sharedRecord={shared}
              onRevoke={handleRevoke}
            />
          ))
        ) : (
          <p>No shared records. Create your first share link above.</p>
        )}
      </div>
    </div>
  );
};

export default Sharing;
