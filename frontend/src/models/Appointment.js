export class Appointment {
  constructor(data = {}) {
    this.id = data.id;
    this.userId = data.user_id;
    this.doctorName = data.doctor_name || '';
    this.specialty = data.specialty || '';
    this.hospital = data.hospital || '';
    this.location = data.location || '';
    this.appointmentDate = data.appointment_date || '';
    this.notes = data.notes || '';
    this.isCompleted = data.is_completed || false;
    this.reminderSent = data.reminder_sent || false;
    this.createdAt = data.created_at;
    this.updatedAt = data.updated_at;
  }

  toJSON() {
    return {
      doctor_name: this.doctorName,
      specialty: this.specialty,
      hospital: this.hospital,
      location: this.location,
      appointment_date: this.appointmentDate,
      notes: this.notes,
      is_completed: this.isCompleted,
    };
  }

  get formattedDate() {
    if (!this.appointmentDate) return '';
    return new Date(this.appointmentDate).toLocaleString();
  }
}

