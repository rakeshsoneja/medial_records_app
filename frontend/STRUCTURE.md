# Frontend Code Structure

This document describes the organized structure of the frontend React application.

## Directory Structure

```
frontend/src/
├── components/          # Reusable UI components
│   ├── cards/          # Card components for displaying records
│   │   ├── AppointmentCard.js
│   │   ├── InsuranceCard.js
│   │   ├── LabReportCard.js
│   │   ├── MedicationCard.js
│   │   ├── PrescriptionCard.js
│   │   ├── ReminderCard.js
│   │   ├── SharedRecordCard.js
│   │   └── Card.css
│   ├── forms/          # Form components for creating/editing records
│   │   ├── AppointmentForm.js
│   │   ├── InsuranceForm.js
│   │   ├── LabReportForm.js
│   │   ├── MedicationForm.js
│   │   ├── PrescriptionForm.js
│   │   ├── ReminderForm.js
│   │   ├── ShareLinkForm.js
│   │   └── Form.css
│   ├── Layout.js       # Main layout component with navigation
│   ├── Layout.css
│   └── PrivateRoute.js # Route protection component
├── controllers/        # Business logic and event handlers
│   ├── AppointmentController.js
│   ├── DashboardController.js
│   ├── InsuranceController.js
│   ├── LabReportController.js
│   ├── MedicationController.js
│   ├── PrescriptionController.js
│   ├── ReminderController.js
│   ├── SharingController.js
│   └── index.js
├── models/            # Data models/types
│   ├── Appointment.js
│   ├── HealthInsurance.js
│   ├── LabReport.js
│   ├── Medication.js
│   ├── Prescription.js
│   ├── Reminder.js
│   ├── SharedRecord.js
│   ├── User.js
│   └── index.js
├── pages/            # Page components (orchestrate components, controllers, models)
│   ├── Appointments.js
│   ├── Dashboard.js
│   ├── Insurance.js
│   ├── LabReports.js
│   ├── Login.js
│   ├── Medications.js
│   ├── Prescriptions.js
│   ├── Register.js
│   ├── Reminders.js
│   ├── Sharing.js
│   └── *.css
├── contexts/         # React contexts for global state
│   └── AuthContext.js
├── services/         # API service layer
│   └── api.js
├── App.js            # Main app component with routing
└── index.js          # Application entry point
```

## Architecture Pattern

### Models (`/models`)
- **Purpose**: Define data structures and business objects
- **Responsibilities**:
  - Data transformation (API response ↔ UI format)
  - Data validation
  - Computed properties (getters)
  - Serialization (toJSON methods)

**Example:**
```javascript
// models/Prescription.js
export class Prescription {
  constructor(data = {}) {
    this.medicineName = data.medicine_name || '';
    // ... map API fields to UI-friendly names
  }
  
  toJSON() {
    return {
      medicine_name: this.medicineName,
      // ... convert back to API format
    };
  }
}
```

### Controllers (`/controllers`)
- **Purpose**: Handle business logic and API interactions
- **Responsibilities**:
  - API calls
  - Data transformation using models
  - Error handling
  - Return standardized response objects

**Example:**
```javascript
// controllers/PrescriptionController.js
async create(prescriptionData) {
  const prescription = new Prescription(prescriptionData);
  const response = await api.post('/prescriptions', prescription.toJSON());
  return {
    success: true,
    data: new Prescription(response.data),
  };
}
```

### Components (`/components`)
- **Purpose**: Reusable UI elements
- **Structure**:
  - **Forms** (`/forms`): Input forms for creating/editing records
  - **Cards** (`/cards`): Display components for showing record data
  - **Layout**: Navigation and page structure
- **Responsibilities**:
  - Present data
  - Handle user interactions
  - Emit events to parent components

**Example:**
```javascript
// components/forms/PrescriptionForm.js
const PrescriptionForm = ({ formData, onChange, onSubmit, onCancel }) => {
  // Form UI and validation
};
```

### Pages (`/pages`)
- **Purpose**: Orchestrate components, controllers, and models
- **Responsibilities**:
  - State management
  - Event handling
  - Coordinate between controllers and components
  - Page-level logic

**Example:**
```javascript
// pages/Prescriptions.js
const Prescriptions = () => {
  const [prescriptions, setPrescriptions] = useState([]);
  
  const fetchPrescriptions = async () => {
    const result = await PrescriptionController.fetchAll();
    if (result.success) {
      setPrescriptions(result.data); // Array of Prescription models
    }
  };
  
  return (
    <>
      <PrescriptionForm onSubmit={handleSubmit} />
      {prescriptions.map(p => <PrescriptionCard prescription={p} />)}
    </>
  );
};
```

## Data Flow

1. **User Action** → Page component
2. **Page** → Controller method
3. **Controller** → API service
4. **API Response** → Controller transforms to Model
5. **Model** → Page state
6. **State** → Components render

## Benefits of This Structure

1. **Separation of Concerns**: Clear boundaries between UI, logic, and data
2. **Reusability**: Components and controllers can be reused across pages
3. **Testability**: Each layer can be tested independently
4. **Maintainability**: Easy to locate and modify specific functionality
5. **Scalability**: Easy to add new features following the same pattern
6. **Type Safety**: Models provide consistent data structures

## Adding New Features

To add a new record type (e.g., "Vehicle Insurance"):

1. **Create Model**: `models/VehicleInsurance.js`
2. **Create Controller**: `controllers/VehicleInsuranceController.js`
3. **Create Form Component**: `components/forms/VehicleInsuranceForm.js`
4. **Create Card Component**: `components/cards/VehicleInsuranceCard.js`
5. **Create Page**: `pages/VehicleInsurance.js`
6. **Add Route**: Update `App.js` with new route

This pattern ensures consistency and makes the codebase easily navigable.




