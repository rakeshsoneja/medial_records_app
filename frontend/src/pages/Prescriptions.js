import React, { useState, useEffect } from 'react';
import PrescriptionController from '../controllers/PrescriptionController';
import PrescriptionForm from '../components/forms/PrescriptionForm';
import PrescriptionCard from '../components/cards/PrescriptionCard';
import './Records.css';

const Prescriptions = () => {
  const [prescriptions, setPrescriptions] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState({
    medicine_name: '',
    dosage: '',
    instructions: '',
    prescribing_doctor: '',
    doctor_specialty: '',
    hospital: '',
    prescription_date: new Date().toISOString().split('T')[0],
  });

  useEffect(() => {
    fetchPrescriptions();
  }, []);

  const fetchPrescriptions = async () => {
    setLoading(true);
    const result = await PrescriptionController.fetchAll();
    if (result.success) {
      setPrescriptions(result.data);
    } else {
      alert(result.error);
    }
    setLoading(false);
  };

  const handleSubmit = async () => {
    const result = await PrescriptionController.create(formData);
    if (result.success) {
      setShowForm(false);
      setFormData({
        medicine_name: '',
        dosage: '',
        instructions: '',
        prescribing_doctor: '',
        doctor_specialty: '',
        hospital: '',
        prescription_date: new Date().toISOString().split('T')[0],
      });
      fetchPrescriptions();
    } else {
      alert(result.error);
    }
  };

  const handleDelete = async (id) => {
    if (window.confirm('Are you sure you want to delete this prescription?')) {
      const result = await PrescriptionController.delete(id);
      if (result.success) {
        fetchPrescriptions();
      } else {
        alert(result.error);
      }
    }
  };

  if (loading) {
    return <div className="loading">Loading prescriptions...</div>;
  }

  return (
    <div className="records-page">
      <div className="page-header">
        <h1>Prescriptions</h1>
        <button className="btn btn-primary" onClick={() => setShowForm(!showForm)}>
          {showForm ? 'Cancel' : 'Add Prescription'}
        </button>
      </div>

      {showForm && (
        <PrescriptionForm
          formData={formData}
          onChange={setFormData}
          onSubmit={handleSubmit}
          onCancel={() => setShowForm(false)}
        />
      )}

      <div className="records-list">
        {prescriptions.length > 0 ? (
          prescriptions.map((prescription) => (
            <PrescriptionCard
              key={prescription.id}
              prescription={prescription}
              onDelete={handleDelete}
            />
          ))
        ) : (
          <p>No prescriptions found. Add your first prescription above.</p>
        )}
      </div>
    </div>
  );
};

export default Prescriptions;
