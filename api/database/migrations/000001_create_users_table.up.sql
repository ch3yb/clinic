-- Users Table
CREATE TABLE users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    first_name TEXT NOT NULL DEFAULT '',  -- Default empty string
    last_name TEXT NOT NULL DEFAULT '',   -- Default empty string
    email TEXT NOT NULL UNIQUE DEFAULT '',-- Default empty string
    phone_number TEXT DEFAULT '',         -- Default empty string
    role TEXT CHECK(role IN ('admin', 'doctor', 'nurse', 'patient')) NOT NULL DEFAULT '',  -- Default empty string
    address TEXT DEFAULT '',              -- Default empty string
    date_of_birth DATE,
    gender TEXT CHECK(gender IN ('male', 'female')) DEFAULT '',  -- Default empty string
    date_joined TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP,
    status TEXT CHECK(status IN ('active', 'inactive', 'suspended')) DEFAULT 'active',
    profile_picture TEXT DEFAULT '',     -- Default empty string
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
    notes TEXT DEFAULT '',  -- Default empty string
    blood_type TEXT CHECK(blood_type IN ('A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-', 'EMPTY')) DEFAULT 'EMPTY',
    emergency_contact_name TEXT DEFAULT '',  -- Default empty string
    emergency_contact_phone TEXT DEFAULT '', -- Default empty string
    insurance_provider TEXT DEFAULT '',      -- Default empty string
    insurance_policy_number TEXT DEFAULT '',-- Default empty string
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    first_name TEXT NOT NULL DEFAULT '',     -- Default empty string
    last_name TEXT NOT NULL DEFAULT '',      -- Default empty string
    email TEXT UNIQUE DEFAULT '',            -- Default empty string
    phone_number TEXT DEFAULT '',            -- Default empty string
    address TEXT DEFAULT '',                 -- Default empty string
    date_of_birth DATE,
    gender TEXT CHECK(gender IN ('m', 'f')) DEFAULT '',  -- Default empty string
    status TEXT CHECK(status IN ('active', 'inactive', 'suspended')) DEFAULT 'active',
    profile_picture TEXT DEFAULT ''          -- Default empty string
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
    medications TEXT DEFAULT '', -- Use TEXT for JSON
    visit_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    instructions TEXT DEFAULT '',  -- Default empty string
    FOREIGN KEY (visit_id) REFERENCES visits(visit_id)
);

-- Prescription Items Table
CREATE TABLE prescription_items (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    medication_name TEXT DEFAULT '',  -- Default empty string
    dosage TEXT DEFAULT '',           -- Default empty string
    frequency TEXT DEFAULT '',        -- Default empty string
    duration INTEGER DEFAULT 0,       -- Default value for integer
    refills INTEGER DEFAULT 0         -- Default value for integer
);

-- Visits Table
CREATE TABLE visits (
    visit_id INTEGER PRIMARY KEY AUTOINCREMENT,
    patient_id INTEGER NOT NULL,
    prescription_id INTEGER,
    details TEXT DEFAULT '',          -- Default empty string
    symptoms TEXT DEFAULT '',         -- Default empty string
    diagnosis TEXT DEFAULT '',        -- Default empty string
    prescription TEXT DEFAULT '',     -- Default empty string
    doctor_inquiry TEXT DEFAULT '',   -- Default empty string
    notes TEXT DEFAULT '',            -- Default empty string
    visit_date TIMESTAMP NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES patients(patient_id) ON DELETE CASCADE,
    FOREIGN KEY (prescription_id) REFERENCES prescriptions(prescription_id) ON DELETE CASCADE
);

-- Diseases Table
CREATE TABLE diseases (
    disease_id INTEGER PRIMARY KEY AUTOINCREMENT,
    disease_name TEXT NOT NULL UNIQUE DEFAULT ''  -- Default empty string
);

-- Patients Diseases Table
CREATE TABLE patients_diseases (
    patient_id INTEGER,
    disease_id INTEGER,
    PRIMARY KEY (patient_id, disease_id),
    FOREIGN KEY (patient_id) REFERENCES patients(patient_id),
    FOREIGN KEY (disease_id) REFERENCES diseases(disease_id)
);
