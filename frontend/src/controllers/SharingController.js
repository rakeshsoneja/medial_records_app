import api from '../services/api';
import { SharedRecord } from '../models';

class SharingController {
  async fetchMyShares() {
    try {
      const response = await api.get('/sharing/my-shares');
      return {
        success: true,
        data: response.data.data.map(item => new SharedRecord(item)),
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to fetch shared records',
      };
    }
  }

  async createShareLink(shareData) {
    try {
      // Parse record IDs if they're a string
      let recordIds = shareData.record_ids;
      if (typeof recordIds === 'string') {
        recordIds = recordIds
          .split(',')
          .map(id => id.trim())
          .filter(id => id);
      }

      const payload = {
        ...shareData,
        record_ids: recordIds,
      };

      const response = await api.post('/sharing/create', payload);
      return {
        success: true,
        data: new SharedRecord(response.data.shared_record),
        shareUrl: response.data.share_url,
      };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to create share link',
      };
    }
  }

  async revokeShareLink(id) {
    try {
      await api.post(`/sharing/${id}/revoke`);
      return { success: true };
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Failed to revoke share link',
      };
    }
  }
}

export default new SharingController();




