import api from '../services/api';
import { HealthInsurance } from '../models';

class InsuranceController {
  async fetchAll() {
    try {
      const response = await api.get('/insurance');
      return {
        success: true,
        data: response.data.data.map(item => new HealthInsurance(item)),
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to fetch insurance records',
      };
    }
  }

  async create(insuranceData) {
    try {
      const insurance = new HealthInsurance(insuranceData);
      const response = await api.post('/insurance', insurance.toJSON());
      return {
        success: true,
        data: new HealthInsurance(response.data),
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to create insurance record',
      };
    }
  }

  async update(id, updates) {
    try {
      const response = await api.put(`/insurance/${id}`, updates);
      return {
        success: true,
        data: response.data,
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to update insurance record',
      };
    }
  }

  async delete(id) {
    try {
      await api.delete(`/insurance/${id}`);
      return { success: true };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to delete insurance record',
      };
    }
  }
}

export default new InsuranceController();

