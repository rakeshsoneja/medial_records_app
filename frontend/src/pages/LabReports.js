import React, { useState, useEffect } from 'react';
import LabReportController from '../controllers/LabReportController';
import LabReportForm from '../components/forms/LabReportForm';
import LabReportCard from '../components/cards/LabReportCard';
import './Records.css';

const LabReports = () => {
  const [labReports, setLabReports] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState({
    test_type: '',
    lab_name: '',
    test_date: new Date().toISOString().split('T')[0],
    report_url: '',
    notes: '',
  });

  useEffect(() => {
    fetchLabReports();
  }, []);

  const fetchLabReports = async () => {
    setLoading(true);
    const result = await LabReportController.fetchAll();
    if (result.success) {
      setLabReports(result.data);
    } else {
      alert(result.error);
    }
    setLoading(false);
  };

  const handleSubmit = async () => {
    const result = await LabReportController.create(formData);
    if (result.success) {
      setShowForm(false);
      setFormData({
        test_type: '',
        lab_name: '',
        test_date: new Date().toISOString().split('T')[0],
        report_url: '',
        notes: '',
      });
      fetchLabReports();
    } else {
      alert(result.error);
    }
  };

  if (loading) {
    return <div className="loading">Loading lab reports...</div>;
  }

  return (
    <div className="records-page">
      <div className="page-header">
        <h1>Lab Reports</h1>
        <button className="btn btn-primary" onClick={() => setShowForm(!showForm)}>
          {showForm ? 'Cancel' : 'Add Lab Report'}
        </button>
      </div>

      {showForm && (
        <LabReportForm
          formData={formData}
          onChange={setFormData}
          onSubmit={handleSubmit}
          onCancel={() => setShowForm(false)}
        />
      )}

      <div className="records-list">
        {labReports.length > 0 ? (
          labReports.map((report) => (
            <LabReportCard key={report.id} labReport={report} />
          ))
        ) : (
          <p>No lab reports found. Add your first lab report above.</p>
        )}
      </div>
    </div>
  );
};

export default LabReports;
