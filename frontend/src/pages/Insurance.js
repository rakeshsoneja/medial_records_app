import React, { useState, useEffect } from 'react';
import InsuranceController from '../controllers/InsuranceController';
import InsuranceForm from '../components/forms/InsuranceForm';
import InsuranceCard from '../components/cards/InsuranceCard';
import './Records.css';

const Insurance = () => {
  const [insurances, setInsurances] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState({
    insurance_provider: '',
    policy_number: '',
    group_number: '',
    member_id: '',
    effective_date: '',
    expiration_date: '',
    notes: '',
  });

  useEffect(() => {
    fetchInsurances();
  }, []);

  const fetchInsurances = async () => {
    setLoading(true);
    const result = await InsuranceController.fetchAll();
    if (result.success) {
      setInsurances(result.data);
    } else {
      alert(result.error);
    }
    setLoading(false);
  };

  const handleSubmit = async () => {
    const result = await InsuranceController.create(formData);
    if (result.success) {
      setShowForm(false);
      setFormData({
        insurance_provider: '',
        policy_number: '',
        group_number: '',
        member_id: '',
        effective_date: '',
        expiration_date: '',
        notes: '',
      });
      fetchInsurances();
    } else {
      alert(result.error);
    }
  };

  if (loading) {
    return <div className="loading">Loading insurance records...</div>;
  }

  return (
    <div className="records-page">
      <div className="page-header">
        <h1>Health Insurance</h1>
        <button className="btn btn-primary" onClick={() => setShowForm(!showForm)}>
          {showForm ? 'Cancel' : 'Add Insurance'}
        </button>
      </div>

      {showForm && (
        <InsuranceForm
          formData={formData}
          onChange={setFormData}
          onSubmit={handleSubmit}
          onCancel={() => setShowForm(false)}
        />
      )}

      <div className="records-list">
        {insurances.length > 0 ? (
          insurances.map((insurance) => (
            <InsuranceCard key={insurance.id} insurance={insurance} />
          ))
        ) : (
          <p>No insurance records found. Add your first insurance record above.</p>
        )}
      </div>
    </div>
  );
};

export default Insurance;
