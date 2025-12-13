import api from '../services/api';
import { Medication } from '../models';

class MedicationController {
  async fetchAll(activeOnly = false) {
    try {
      const response = await api.get('/medications', {
        params: { active: activeOnly },
      });
      return {
        success: true,
        data: response.data.data.map(item => new Medication(item)),
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to fetch medications',
      };
    }
  }

  async fetchNeedingRefill() {
    try {
      const response = await api.get('/medications/refill-needed');
      return {
        success: true,
        data: response.data.data.map(item => new Medication(item)),
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to fetch medications needing refill',
      };
    }
  }

  async create(medicationData) {
    try {
      const medication = new Medication(medicationData);
      const response = await api.post('/medications', medication.toJSON());
      return {
        success: true,
        data: new Medication(response.data),
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to create medication',
      };
    }
  }

  async update(id, updates) {
    try {
      const response = await api.put(`/medications/${id}`, updates);
      return {
        success: true,
        data: response.data,
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to update medication',
      };
    }
  }

  async delete(id) {
    try {
      await api.delete(`/medications/${id}`);
      return { success: true };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to delete medication',
      };
    }
  }
}

export default new MedicationController();

