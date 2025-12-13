export class Prescription {
  constructor(data = {}) {
    this.id = data.id;
    this.userId = data.user_id;
    this.medicineName = data.medicine_name || '';
    this.dosage = data.dosage || '';
    this.instructions = data.instructions || '';
    this.prescribingDoctor = data.prescribing_doctor || '';
    this.doctorSpecialty = data.doctor_specialty || '';
    this.hospital = data.hospital || '';
    this.prescriptionDate = data.prescription_date || new Date().toISOString().split('T')[0];
    this.attachmentUrl = data.attachment_url;
    this.attachmentType = data.attachment_type;
    this.isActive = data.is_active !== undefined ? data.is_active : true;
    this.createdAt = data.created_at;
    this.updatedAt = data.updated_at;
  }

  toJSON() {
    return {
      medicine_name: this.medicineName,
      dosage: this.dosage,
      instructions: this.instructions,
      prescribing_doctor: this.prescribingDoctor,
      doctor_specialty: this.doctorSpecialty,
      hospital: this.hospital,
      prescription_date: this.prescriptionDate,
      attachment_url: this.attachmentUrl,
      attachment_type: this.attachmentType,
      is_active: this.isActive,
    };
  }
}

