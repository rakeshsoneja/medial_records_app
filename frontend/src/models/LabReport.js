export class LabReport {
  constructor(data = {}) {
    this.id = data.id;
    this.userId = data.user_id;
    this.testType = data.test_type || '';
    this.labName = data.lab_name || '';
    this.testDate = data.test_date || new Date().toISOString().split('T')[0];
    this.reportUrl = data.report_url || '';
    this.reportType = data.report_type;
    this.notes = data.notes || '';
    this.createdAt = data.created_at;
    this.updatedAt = data.updated_at;
  }

  toJSON() {
    return {
      test_type: this.testType,
      lab_name: this.labName,
      test_date: this.testDate,
      report_url: this.reportUrl,
      report_type: this.reportType,
      notes: this.notes,
    };
  }

  get formattedDate() {
    if (!this.testDate) return '';
    return new Date(this.testDate).toLocaleDateString();
  }
}




