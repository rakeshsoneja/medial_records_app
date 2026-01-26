import api from '../services/api';
import { LabReport } from '../models';

class LabReportController {
  async fetchAll(limit = 10, offset = 0) {
    try {
      const response = await api.get('/lab-reports', {
        params: { limit, offset },
      });
      return {
        success: true,
        data: response.data.data.map(item => new LabReport(item)),
        total: response.data.total,
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to fetch lab reports',
      };
    }
  }

  async create(labReportData) {
    try {
      const labReport = new LabReport(labReportData);
      const response = await api.post('/lab-reports', labReport.toJSON());
      return {
        success: true,
        data: new LabReport(response.data),
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to create lab report',
      };
    }
  }

  async update(id, updates) {
    try {
      const response = await api.put(`/lab-reports/${id}`, updates);
      return {
        success: true,
        data: response.data,
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to update lab report',
      };
    }
  }

  async delete(id) {
    try {
      await api.delete(`/lab-reports/${id}`);
      return { success: true };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to delete lab report',
      };
    }
  }
}

export default new LabReportController();




