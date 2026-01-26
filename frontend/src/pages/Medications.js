import React, { useState, useEffect } from 'react';
import MedicationController from '../controllers/MedicationController';
import MedicationForm from '../components/forms/MedicationForm';
import MedicationCard from '../components/cards/MedicationCard';
import './Records.css';

const Medications = () => {
  const [medications, setMedications] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState({
    medicine_name: '',
    dosage: '',
    frequency: '',
    pharmacy_name: '',
    pharmacy_phone: '',
    pharmacy_address: '',
    refill_reminder_days: 7,
  });

  useEffect(() => {
    fetchMedications();
  }, []);

  const fetchMedications = async () => {
    setLoading(true);
    const result = await MedicationController.fetchAll(true);
    if (result.success) {
      setMedications(result.data);
    } else {
      alert(result.error);
    }
    setLoading(false);
  };

  const handleSubmit = async () => {
    const result = await MedicationController.create(formData);
    if (result.success) {
      setShowForm(false);
      setFormData({
        medicine_name: '',
        dosage: '',
        frequency: '',
        pharmacy_name: '',
        pharmacy_phone: '',
        pharmacy_address: '',
        refill_reminder_days: 7,
      });
      fetchMedications();
    } else {
      alert(result.error);
    }
  };

  if (loading) {
    return <div className="loading">Loading medications...</div>;
  }

  return (
    <div className="records-page">
      <div className="page-header">
        <h1>Medications</h1>
        <button className="btn btn-primary" onClick={() => setShowForm(!showForm)}>
          {showForm ? 'Cancel' : 'Add Medication'}
        </button>
      </div>

      {showForm && (
        <MedicationForm
          formData={formData}
          onChange={setFormData}
          onSubmit={handleSubmit}
          onCancel={() => setShowForm(false)}
        />
      )}

      <div className="records-list">
        {medications.length > 0 ? (
          medications.map((medication) => (
            <MedicationCard key={medication.id} medication={medication} />
          ))
        ) : (
          <p>No medications found. Add your first medication above.</p>
        )}
      </div>
      <div className="fixed-bottom-bar">
        <button className="btn btn-primary" onClick={() => setShowForm(!showForm)}>
          {showForm ? 'Cancel' : 'Add Medication'}
        </button>
      </div>
    </div>
  );
};

export default Medications;
