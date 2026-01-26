export class HealthInsurance {
  constructor(data = {}) {
    this.id = data.id;
    this.userId = data.user_id;
    this.insuranceProvider = data.insurance_provider || '';
    this.policyNumber = data.policy_number || '';
    this.groupNumber = data.group_number || '';
    this.memberId = data.member_id || '';
    this.effectiveDate = data.effective_date;
    this.expirationDate = data.expiration_date;
    this.notes = data.notes || '';
    this.createdAt = data.created_at;
    this.updatedAt = data.updated_at;
  }

  toJSON() {
    return {
      insurance_provider: this.insuranceProvider,
      policy_number: this.policyNumber,
      group_number: this.groupNumber,
      member_id: this.memberId,
      effective_date: this.effectiveDate,
      expiration_date: this.expirationDate,
      notes: this.notes,
    };
  }
}




