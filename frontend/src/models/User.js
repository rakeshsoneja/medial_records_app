export class User {
  constructor(data) {
    this.id = data.id;
    this.email = data.email;
    this.phone = data.phone;
    this.firstName = data.first_name || data.firstName;
    this.lastName = data.last_name || data.lastName;
    this.dateOfBirth = data.date_of_birth;
    this.role = data.role || 'patient';
    this.isEmailVerified = data.is_email_verified;
    this.isPhoneVerified = data.is_phone_verified;
  }

  get fullName() {
    return `${this.firstName} ${this.lastName}`;
  }
}

