export class Medication {
  constructor(data = {}) {
    this.id = data.id;
    this.userId = data.user_id;
    this.medicineName = data.medicine_name || '';
    this.dosage = data.dosage || '';
    this.frequency = data.frequency || '';
    this.pharmacyName = data.pharmacy_name || '';
    this.pharmacyPhone = data.pharmacy_phone || '';
    this.pharmacyAddress = data.pharmacy_address || '';
    this.lastRefillDate = data.last_refill_date;
    this.nextRefillDate = data.next_refill_date;
    this.refillReminderDays = data.refill_reminder_days || 7;
    this.isActive = data.is_active !== undefined ? data.is_active : true;
    this.createdAt = data.created_at;
    this.updatedAt = data.updated_at;
  }

  toJSON() {
    return {
      medicine_name: this.medicineName,
      dosage: this.dosage,
      frequency: this.frequency,
      pharmacy_name: this.pharmacyName,
      pharmacy_phone: this.pharmacyPhone,
      pharmacy_address: this.pharmacyAddress,
      last_refill_date: this.lastRefillDate,
      next_refill_date: this.nextRefillDate,
      refill_reminder_days: this.refillReminderDays,
      is_active: this.isActive,
    };
  }
}

