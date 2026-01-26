import api from '../services/api';
import { Prescription } from '../models';

class PrescriptionController {
  async fetchAll(limit = 10, offset = 0) {
    try {
      const response = await api.get('/prescriptions', {
        params: { limit, offset },
      });
      return {
        success: true,
        data: response.data.data.map(item => new Prescription(item)),
        total: response.data.total,
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to fetch prescriptions',
      };
    }
  }

  async fetchById(id) {
    try {
      const response = await api.get(`/prescriptions/${id}`);
      return {
        success: true,
        data: new Prescription(response.data),
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to fetch prescription',
      };
    }
  }

  async create(prescriptionData) {
    try {
      const prescription = new Prescription(prescriptionData);
      const response = await api.post('/prescriptions', prescription.toJSON());
      return {
        success: true,
        data: new Prescription(response.data),
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to create prescription',
      };
    }
  }

  async update(id, updates) {
    try {
      const response = await api.put(`/prescriptions/${id}`, updates);
      return {
        success: true,
        data: response.data,
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to update prescription',
      };
    }
  }

  async delete(id) {
    try {
      await api.delete(`/prescriptions/${id}`);
      return { success: true };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to delete prescription',
      };
    }
  }
}

export default new PrescriptionController();




