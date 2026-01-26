import api from '../services/api';
import { Reminder } from '../models';

class ReminderController {
  async fetchAll(upcomingOnly = false) {
    try {
      const response = await api.get('/reminders', {
        params: { upcoming: upcomingOnly },
      });
      return {
        success: true,
        data: response.data.data.map(item => new Reminder(item)),
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to fetch reminders',
      };
    }
  }

  async fetchUpcoming(days = 30) {
    try {
      const response = await api.get('/reminders/upcoming', {
        params: { days },
      });
      return {
        success: true,
        data: response.data.data.map(item => new Reminder(item)),
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to fetch upcoming reminders',
      };
    }
  }

  async create(reminderData) {
    try {
      const reminder = new Reminder(reminderData);
      const response = await api.post('/reminders', reminder.toJSON());
      return {
        success: true,
        data: new Reminder(response.data),
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to create reminder',
      };
    }
  }

  async update(id, updates) {
    try {
      const response = await api.put(`/reminders/${id}`, updates);
      return {
        success: true,
        data: response.data,
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to update reminder',
      };
    }
  }

  async delete(id) {
    try {
      await api.delete(`/reminders/${id}`);
      return { success: true };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to delete reminder',
      };
    }
  }
}

export default new ReminderController();




