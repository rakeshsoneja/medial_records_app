import api from '../services/api';
import { Appointment } from '../models';

class AppointmentController {
  async fetchAll(limit = 10, offset = 0, upcomingOnly = false) {
    try {
      const response = await api.get('/appointments', {
        params: { limit, offset, upcoming: upcomingOnly },
      });
      return {
        success: true,
        data: response.data.data.map(item => new Appointment(item)),
        total: response.data.total,
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to fetch appointments',
      };
    }
  }

  async create(appointmentData) {
    try {
      const appointment = new Appointment(appointmentData);
      const response = await api.post('/appointments', appointment.toJSON());
      return {
        success: true,
        data: new Appointment(response.data),
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to create appointment',
      };
    }
  }

  async update(id, updates) {
    try {
      const response = await api.put(`/appointments/${id}`, updates);
      return {
        success: true,
        data: response.data,
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to update appointment',
      };
    }
  }

  async delete(id) {
    try {
      await api.delete(`/appointments/${id}`);
      return { success: true };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to delete appointment',
      };
    }
  }
}

export default new AppointmentController();




