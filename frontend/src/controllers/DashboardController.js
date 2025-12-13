import api from '../services/api';
import { Prescription, Appointment, LabReport, Medication, Reminder } from '../models';

class DashboardController {
  async fetchDashboard() {
    try {
      const response = await api.get('/dashboard');
      return {
        success: true,
        data: {
          prescriptions: response.data.prescriptions.map(item => new Prescription(item)),
          appointments: response.data.appointments.map(item => new Appointment(item)),
          labReports: response.data.lab_reports.map(item => new LabReport(item)),
          medications: response.data.medications.map(item => new Medication(item)),
          reminders: response.data.reminders.map(item => new Reminder(item)),
        },
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to fetch dashboard data',
      };
    }
  }
}

export default new DashboardController();

