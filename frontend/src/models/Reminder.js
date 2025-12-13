export class Reminder {
  constructor(data = {}) {
    this.id = data.id;
    this.userId = data.user_id;
    this.title = data.title || '';
    this.description = data.description || '';
    this.reminderDate = data.reminder_date || '';
    this.reminderType = data.reminder_type || 'checkup';
    this.isCompleted = data.is_completed || false;
    this.isRecurring = data.is_recurring || false;
    this.recurrenceInterval = data.recurrence_interval || '';
    this.createdAt = data.created_at;
    this.updatedAt = data.updated_at;
  }

  toJSON() {
    return {
      title: this.title,
      description: this.description,
      reminder_date: this.reminderDate,
      reminder_type: this.reminderType,
      is_completed: this.isCompleted,
      is_recurring: this.isRecurring,
      recurrence_interval: this.recurrenceInterval,
    };
  }

  get formattedDate() {
    if (!this.reminderDate) return '';
    return new Date(this.reminderDate).toLocaleString();
  }
}

