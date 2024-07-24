-- Users Table
CREATE TABLE users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL ,   -- Default empty string
    email TEXT NOT NULL UNIQUE,
    phone_number TEXT DEFAULT '',         -- Default empty string
    role TEXT CHECK(role IN ('admin', 'doctor', 'nurse', 'patient')) NOT NULL DEFAULT '',  -- Default empty string
    address TEXT,
    date_of_birth DATE,
    gender TEXT CHECK(gender IN ('male', 'female')) DEFAULT '',  -- Default empty string
    date_joined TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP,
    status TEXT CHECK(status IN ('active', 'inactive', 'suspended')) DEFAULT 'active',
    profile_picture TEXT DEFAULT '',     -- Default empty string
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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
    emergency_contact_name TEXT,
    emergency_contact_phone TEXT,
    insurance_provider TEXT,
    insurance_policy_number TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE ,
    phone_number TEXT,
    address TEXT,
    date_of_birth DATE,
    gender TEXT CHECK(gender IN ('m', 'f')) DEFAULT 'm',
    status TEXT CHECK(status IN ('active', 'inactive', 'suspended')) DEFAULT 'active',
    profile_picture TEXT DEFAULT '',
    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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
-- Users Table
INSERT INTO users (username, password, first_name, last_name, email, phone_number, role, address, date_of_birth, gender, date_joined, last_login, status, profile_picture) VALUES
('john_doe', 'password123', 'John', 'Doe', 'john.doe@example.com', '1234567890', 'admin', '123 Elm Street', '1980-01-01', 'male', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'active', ''),
('jane_smith', 'password123', 'Jane', 'Smith', 'jane.smith@example.com', '2345678901', 'doctor', '456 Oak Avenue', '1985-02-02', 'female', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'active', ''),
('alice_jones', 'password123', 'Alice', 'Jones', 'alice.jones@example.com', '3456789012', 'nurse', '789 Pine Road', '1990-03-03', 'female', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'active', ''),
('bob_brown', 'password123', 'Bob', 'Brown', 'bob.brown@example.com', '4567890123', 'patient', '321 Birch Boulevard', '1995-04-04', 'male', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'active', ''),
('charlie_davis', 'password123', 'Charlie', 'Davis', 'charlie.davis@example.com', '5678901234', 'patient', '654 Maple Lane', '2000-05-05', 'male', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'active', ''),
('david_lee', 'password123', 'David', 'Lee', 'david.lee@example.com', '6789012345', 'doctor', '987 Cedar Street', '1975-06-06', 'male', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'active', ''),
('emily_clark', 'password123', 'Emily', 'Clark', 'emily.clark@example.com', '7890123456', 'nurse', '123 Fir Avenue', '1988-07-07', 'female', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'active', ''),
('frank_miller', 'password123', 'Frank', 'Miller', 'frank.miller@example.com', '8901234567', 'admin', '456 Spruce Road', '1979-08-08', 'male', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'active', ''),
('grace_wilson', 'password123', 'Grace', 'Wilson', 'grace.wilson@example.com', '9012345678', 'patient', '789 Ash Boulevard', '1992-09-09', 'female', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'active', ''),
('hannah_taylor', 'password123', 'Hannah', 'Taylor', 'hannah.taylor@example.com', '0123456789', 'doctor', '321 Redwood Lane', '1983-10-10', 'female', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'active', '');

-- Patients Table
INSERT INTO patients (first_name, last_name, email, phone_number, address, date_of_birth, gender, status, profile_picture, notes, blood_type, emergency_contact_name, emergency_contact_phone, insurance_provider, insurance_policy_number) VALUES
('John', 'Doe', 'john.doe@example.com', '1234567890', '123 Elm Street', '1980-01-01', 'm', 'active', '', '', 'A+', 'Jane Doe', '1234567890', 'Blue Cross', '12345'),
('Jane', 'Smith', 'jane.smith@example.com', '2345678901', '456 Oak Avenue', '1985-02-02', 'f', 'active', '', '', 'B+', 'John Smith', '2345678901', 'UnitedHealth', '67890'),
('Alice', 'Jones', 'alice.jones@example.com', '3456789012', '789 Pine Road', '1990-03-03', 'f', 'active', '', '', 'AB+', 'Bob Jones', '3456789012', 'Aetna', '11223'),
('Bob', 'Brown', 'bob.brown@example.com', '4567890123', '321 Birch Boulevard', '1995-04-04', 'm', 'active', '', '', 'O+', 'Alice Brown', '4567890123', 'Cigna', '44556'),
('Charlie', 'Davis', 'charlie.davis@example.com', '5678901234', '654 Maple Lane', '2000-05-05', 'm', 'active', '', '', 'A-', 'David Davis', '5678901234', 'Humana', '77889'),
('David', 'Lee', 'david.lee@example.com', '6789012345', '987 Cedar Street', '1975-06-06', 'm', 'active', '', '', 'B-', 'Grace Lee', '6789012345', 'Kaiser', '99001'),
('Emily', 'Clark', 'emily.clark@example.com', '7890123456', '123 Fir Avenue', '1988-07-07', 'f', 'active', '', '', 'AB-', 'Hannah Clark', '7890123456', 'Blue Cross', '22334'),
('Frank', 'Miller', 'frank.miller@example.com', '8901234567', '456 Spruce Road', '1979-08-08', 'm', 'active', '', '', 'O-', 'Emily Miller', '8901234567', 'UnitedHealth', '55667'),
('Grace', 'Wilson', 'grace.wilson@example.com', '9012345678', '789 Ash Boulevard', '1992-09-09', 'f', 'active', '', '', 'A+', 'Frank Wilson', '9012345678', 'Aetna', '88900'),
('Hannah', 'Taylor', 'hannah.taylor@example.com', '0123456789', '321 Redwood Lane', '1983-10-10', 'f', 'active', '', '', 'B+', 'John Taylor', '0123456789', 'Cigna', '33221');

-- Prescriptions Table
INSERT INTO prescriptions (medications, visit_id, instructions) VALUES
('{"medication":"Aspirin", "dosage":"100mg"}', 1, 'Take one tablet daily'),
('{"medication":"Ibuprofen", "dosage":"200mg"}', 2, 'Take two tablets after meal'),
('{"medication":"Amoxicillin", "dosage":"500mg"}', 3, 'Take one capsule every 8 hours'),
('{"medication":"Lisinopril", "dosage":"10mg"}', 4, 'Take one tablet every morning'),
('{"medication":"Metformin", "dosage":"500mg"}', 5, 'Take one tablet twice daily'),
('{"medication":"Amlodipine", "dosage":"5mg"}', 6, 'Take one tablet daily'),
('{"medication":"Atorvastatin", "dosage":"20mg"}', 7, 'Take one tablet every night'),
('{"medication":"Omeprazole", "dosage":"20mg"}', 8, 'Take one capsule before breakfast'),
('{"medication":"Albuterol", "dosage":"2 puffs"}', 9, 'Use inhaler as needed'),
('{"medication":"Levothyroxine", "dosage":"50mcg"}', 10, 'Take one tablet daily');

-- Prescription Items Table
INSERT INTO prescription_items (medication_name, dosage, frequency, duration, refills) VALUES
('Aspirin', '100mg', 'daily', 30, 2),
('Ibuprofen', '200mg', 'after meal', 14, 1),
('Amoxicillin', '500mg', 'every 8 hours', 10, 0),
('Lisinopril', '10mg', 'every morning', 30, 3),
('Metformin', '500mg', 'twice daily', 30, 1),
('Amlodipine', '5mg', 'daily', 30, 2),
('Atorvastatin', '20mg', 'every night', 30, 3),
('Omeprazole', '20mg', 'before breakfast', 14, 0),
('Albuterol', '2 puffs', 'as needed', 30, 2),
('Levothyroxine', '50mcg', 'daily', 30, 3);

-- Visits Table
INSERT INTO visits (patient_id, prescription_id, details, symptoms, diagnosis, prescription, doctor_inquiry, notes, visit_date) VALUES
(1, 1, 'Routine check-up', 'None', 'Healthy', '', '', '', CURRENT_TIMESTAMP),
(2, 2, 'Follow-up', 'Headache', 'Migraine', '', '', '', CURRENT_TIMESTAMP),
(3, 3, 'Initial consultation', 'Cough', 'Bronchitis', '', '', '', CURRENT_TIMESTAMP),
(4, 4, 'Follow-up', 'High blood pressure', 'Hypertension', '', '', '', CURRENT_TIMESTAMP),
(5, 5, 'Routine check-up', 'None', 'Healthy', '', '', '', CURRENT_TIMESTAMP),
(6, 6, 'Follow-up', 'Diabetes', 'Diabetes Type 2', '', '', '', CURRENT_TIMESTAMP),
(7, 7, 'Routine check-up', 'None', 'Healthy', '', '', '', CURRENT_TIMESTAMP),
(8, 8, 'Initial consultation', 'Stomach pain', 'Gastritis', '', '', '', CURRENT_TIMESTAMP),
(9, 9, 'Follow-up', 'Asthma', 'Asthma', '', '', '', CURRENT_TIMESTAMP),
(10, 10, 'Routine check-up', 'None', 'Healthy', '', '', '', CURRENT_TIMESTAMP);

-- Diseases Table
INSERT INTO diseases (disease_name) VALUES
('Hypertension'),
('Diabetes Type 2'),
('Asthma'),
('Bronchitis'),
('Migraine'),
('Gastritis'),
('Hyperlipidemia'),
('Anemia'),
('Hypothyroidism'),
('Arthritis');

-- Patients Diseases Table
INSERT INTO patients_diseases (patient_id, disease_id) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5),
(6, 6),
(7, 7),
(8, 8),
(9, 9),
(10, 10);

