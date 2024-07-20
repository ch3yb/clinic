-- Users Table
CREATE TABLE users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    phone_number TEXT,
    role TEXT CHECK(role IN ('admin', 'doctor', 'nurse', 'patient')) NOT NULL,
    address TEXT,
    date_of_birth DATE,
    gender TEXT CHECK(gender IN ('male', 'female')),
    date_joined TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP,
    status TEXT CHECK(status IN ('active', 'inactive', 'suspended')) DEFAULT 'active',
    profile_picture TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Trigger to update 'updated_at' column in users table
CREATE TRIGGER update_users_updated_at
AFTER UPDATE ON users
FOR EACH ROW
BEGIN
    UPDATE users SET updated_at = CURRENT_TIMESTAMP WHERE user_id = old.user_id;
END;

-- Patients Table
CREATE TABLE patients (
    patient_id INTEGER PRIMARY KEY AUTOINCREMENT,
    notes TEXT,
    blood_type TEXT CHECK(blood_type IN ('A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-','EMPTY')) DEFAULT 'EMPTY',
    emergency_contact_name TEXT,
    emergency_contact_phone TEXT,
    insurance_provider TEXT,
    insurance_policy_number TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     first_name TEXT NOT NULL,
        last_name TEXT NOT NULL,
        email TEXT UNIQUE,
        phone_number TEXT,
        address TEXT,
        date_of_birth DATE,
        gender TEXT CHECK(gender IN ('m', 'f')),
        status TEXT CHECK(status IN ('active', 'inactive', 'suspended')) DEFAULT 'active',
        profile_picture TEXT
);

-- Trigger to update 'updated_at' column in patients table
CREATE TRIGGER update_patients_updated_at
AFTER UPDATE ON patients
FOR EACH ROW
BEGIN
    UPDATE patients SET updated_at = CURRENT_TIMESTAMP WHERE patient_id = old.patient_id;
END;

-- Prescriptions Table
CREATE TABLE prescriptions (
    prescription_id INTEGER PRIMARY KEY AUTOINCREMENT,
    medications TEXT, -- SQLite does not have a JSON type, use TEXT for JSON
    visit_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    instructions TEXT,
    FOREIGN KEY (visit_id) REFERENCES visits(visit_id)
);

-- Prescription Items Table
CREATE TABLE prescription_items (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    medication_name TEXT,
    dosage TEXT,
    frequency TEXT,  -- How often to take the medication (e.g., daily, twice a day)
    duration INTEGER,           -- Duration of the prescription (e.g., in days)
    refills INTEGER             -- Number of allowed refills
);

-- Visits Table
CREATE TABLE visits (
    visit_id INTEGER PRIMARY KEY AUTOINCREMENT,
    patient_id INTEGER NOT NULL,
    prescription_id INTEGER,
    details TEXT,
    symptoms TEXT,
    diagnosis TEXT,
    prescription TEXT,
    doctor_inquiry TEXT,
    notes TEXT,
    visit_date TIMESTAMP NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES patients(patient_id) ON DELETE CASCADE,
    FOREIGN KEY (prescription_id) REFERENCES prescriptions(prescription_id) ON DELETE CASCADE
);

-- Diseases Table
CREATE TABLE diseases (
    disease_id INTEGER PRIMARY KEY AUTOINCREMENT,
    disease_name TEXT NOT NULL UNIQUE
);

-- Patients Diseases Table
CREATE TABLE patients_diseases (
    patient_id INTEGER,
    disease_id INTEGER,
    PRIMARY KEY (patient_id, disease_id),
    FOREIGN KEY (patient_id) REFERENCES patients(patient_id),
    FOREIGN KEY (disease_id) REFERENCES diseases(disease_id)
);
