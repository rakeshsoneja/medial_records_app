export class SharedRecord {
  constructor(data = {}) {
    this.id = data.id;
    this.userId = data.user_id;
    this.shareToken = data.share_token;
    this.recordType = data.record_type;
    this.recordIds = data.record_ids;
    this.expiresAt = data.expires_at;
    this.maxAccessCount = data.max_access_count || 0;
    this.currentAccessCount = data.current_access_count || 0;
    this.allowDownload = data.allow_download || false;
    this.recipientEmail = data.recipient_email;
    this.recipientPhone = data.recipient_phone;
    this.shareMethod = data.share_method;
    this.isActive = data.is_active;
    this.createdAt = data.created_at;
    this.updatedAt = data.updated_at;
  }

  get shareUrl() {
    if (!this.shareToken) return '';
    return `${window.location.origin}/share/${this.shareToken}`;
  }

  get isExpired() {
    if (!this.expiresAt) return false;
    return new Date(this.expiresAt) < new Date();
  }

  get accessLimitReached() {
    if (this.maxAccessCount === 0) return false;
    return this.currentAccessCount >= this.maxAccessCount;
  }
}

