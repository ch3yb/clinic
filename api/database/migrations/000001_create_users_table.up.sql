-- Users Table
CREATE TABLE users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,   -- Default empty string
    email TEXT NOT NULL UNIQUE,
    phone_number TEXT DEFAULT '',         -- Default empty string
    role TEXT CHECK(role IN ('admin', 'doctor', 'nurse', 'patient')) NOT NULL DEFAULT '',  -- Default empty string
    address TEXT,
    date_of_birth BIGINT,
    gender TEXT CHECK(gender IN ('male', 'female')) DEFAULT '',  -- Default empty string
    date_joined BIGINT DEFAULT (strftime('%s', 'now')),
    last_login BIGINT,
    status TEXT CHECK(status IN ('active', 'inactive', 'suspended')) DEFAULT 'active',
    profile_picture TEXT DEFAULT '',     -- Default empty string
    created_at BIGINT DEFAULT (strftime('%s', 'now')),
    updated_at BIGINT DEFAULT (strftime('%s', 'now')),
    deleted_at BIGINT DEFAULT NULL
);

-- Trigger to update 'updated_at' column in users table
CREATE TRIGGER update_users_updated_at
AFTER UPDATE ON users
FOR EACH ROW
BEGIN
    UPDATE users SET updated_at = strftime('%s', 'now') WHERE user_id = old.user_id;
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
    created_at BIGINT DEFAULT (strftime('%s', 'now')),
    updated_at BIGINT DEFAULT (strftime('%s', 'now')),
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE,
    phone_number TEXT,
    address TEXT,
    date_of_birth BIGINT,
    gender TEXT CHECK(gender IN ('m', 'f')) DEFAULT 'm',
    status TEXT CHECK(status IN ('active', 'inactive', 'suspended')) DEFAULT 'active',
    profile_picture TEXT DEFAULT '',
    deleted_at BIGINT DEFAULT (strftime('%s', 'now'))
);

-- Trigger to update 'updated_at' column in patients table
CREATE TRIGGER update_patients_updated_at
AFTER UPDATE ON patients
FOR EACH ROW
BEGIN
    UPDATE patients SET updated_at = strftime('%s', 'now') WHERE patient_id = old.patient_id;
END;

-- Visits Table
CREATE TABLE visits (
    visit_id INTEGER PRIMARY KEY AUTOINCREMENT,
    patient_id INTEGER NOT NULL,
    prescription_id INTEGER DEFAULT NULL,
    details TEXT DEFAULT '',          -- Default empty string
    diagnosis TEXT DEFAULT '',        -- Default empty string
    notes TEXT DEFAULT '',            -- Default empty string
    visit_date BIGINT NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES patients(patient_id) ON DELETE CASCADE,
    FOREIGN KEY (prescription_id) REFERENCES prescriptions(prescription_id) ON DELETE CASCADE
);


-- Prescriptions Table
CREATE TABLE prescriptions (
    prescription_id INTEGER PRIMARY KEY AUTOINCREMENT,
    medications TEXT DEFAULT '', -- Use TEXT for JSON
    visit_id INTEGER NOT NULL,
    created_at BIGINT DEFAULT (strftime('%s', 'now')),
    instructions TEXT DEFAULT '',  -- Default empty string
    FOREIGN KEY (visit_id) REFERENCES visits(visit_id)
);

-- Prescription Items Table
CREATE TABLE medicament (
    id SERIAL PRIMARY KEY,
    numero VARCHAR(255) NOT NULL,
    numero_enregistrement VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL,
    denomination_commune_internationale VARCHAR(255),
    nom_de_marque VARCHAR(255),
    forme VARCHAR(255),
    dosage VARCHAR(255),
    cond VARCHAR(255),
    liste TEXT,
    p1 TEXT,
    p2 TEXT,
    obs TEXT,
    laboratoires_detenteur_de_la_decision_enregistrement VARCHAR(255),
    pays_du_laboratoire_detenteur_de_la_decision_enregistrement VARCHAR(255),
    date_enregistrement_initial DATE,
    date_enregistrement_final DATE,
    type VARCHAR(255),
    statut VARCHAR(255),
    duree_de_stabilite VARCHAR(255)
);


-- Prescription Items Table
CREATE TABLE prescription_items (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    prescription_id INTEGER,
    medicament_id INTEGER,
    dosage TEXT DEFAULT '',           -- Default empty string
    frequency TEXT DEFAULT '',        -- Default empty string
    duration INTEGER DEFAULT 0,       -- Default value for integer
    refills INTEGER DEFAULT 0   ,     -- Default value for integer
    FOREIGN KEY (prescription_id) REFERENCES prescriptions(prescription_id),
    FOREIGN KEY (medicament_id) REFERENCES medicament(id)
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

-- Users Table Data Insertion
INSERT INTO users (username, password, first_name, last_name, email, phone_number, role, address, date_of_birth, gender, date_joined, last_login, status, profile_picture) VALUES
('john_doe', 'password123', 'John', 'Doe', 'john.doe@example.com', '1234567890', 'admin', '123 Elm Street', strftime('%s', '1980-01-01'), 'male', strftime('%s', 'now'), strftime('%s', 'now'), 'active', ''),
('jane_smith', 'password123', 'Jane', 'Smith', 'jane.smith@example.com', '2345678901', 'doctor', '456 Oak Avenue', strftime('%s', '1985-02-02'), 'female', strftime('%s', 'now'), strftime('%s', 'now'), 'active', ''),
('alice_jones', 'password123', 'Alice', 'Jones', 'alice.jones@example.com', '3456789012', 'nurse', '789 Pine Road', strftime('%s', '1990-03-03'), 'female', strftime('%s', 'now'), strftime('%s', 'now'), 'active', ''),
('bob_brown', 'password123', 'Bob', 'Brown', 'bob.brown@example.com', '4567890123', 'patient', '321 Birch Boulevard', strftime('%s', '1995-04-04'), 'male', strftime('%s', 'now'), strftime('%s', 'now'), 'active', ''),
('charlie_davis', 'password123', 'Charlie', 'Davis', 'charlie.davis@example.com', '5678901234', 'patient', '654 Maple Lane', strftime('%s', '2000-05-05'), 'male', strftime('%s', 'now'), strftime('%s', 'now'), 'active', ''),
('david_lee', 'password123', 'David', 'Lee', 'david.lee@example.com', '6789012345', 'doctor', '987 Cedar Street', strftime('%s', '1975-06-06'), 'male', strftime('%s', 'now'), strftime('%s', 'now'), 'active', ''),
('emily_clark', 'password123', 'Emily', 'Clark', 'emily.clark@example.com', '7890123456', 'nurse', '123 Fir Avenue', strftime('%s', '1988-07-07'), 'female', strftime('%s', 'now'), strftime('%s', 'now'), 'active', ''),
('frank_miller', 'password123', 'Frank', 'Miller', 'frank.miller@example.com', '8901234567', 'admin', '456 Spruce Road', strftime('%s', '1979-08-08'), 'male', strftime('%s', 'now'), strftime('%s', 'now'), 'active', ''),
('grace_wilson', 'password123', 'Grace', 'Wilson', 'grace.wilson@example.com', '9012345678', 'patient', '789 Ash Boulevard', strftime('%s', '1992-09-09'), 'female', strftime('%s', 'now'), strftime('%s', 'now'), 'active', ''),
('hannah_taylor', 'password123', 'Hannah', 'Taylor', 'hannah.taylor@example.com', '0123456789', 'doctor', '321 Redwood Lane', strftime('%s', '1983-10-10'), 'female', strftime('%s', 'now'), strftime('%s', 'now'), 'active', '');

-- Patients Table Data Insertion
INSERT INTO patients (first_name, last_name, email, phone_number, address, date_of_birth, gender, status, profile_picture, notes, blood_type, emergency_contact_name, emergency_contact_phone, insurance_provider, insurance_policy_number) VALUES
('John', 'Doe', 'john.doe@example.com', '1234567890', '123 Elm Street', strftime('%s', '1980-01-01'), 'm', 'active', '', '', 'A+', 'Jane Doe', '1234567890', 'Blue Cross', '12345'),
('Jane', 'Smith', 'jane.smith@example.com', '2345678901', '456 Oak Avenue', strftime('%s', '1985-02-02'), 'f', 'active', '', '', 'B+', 'John Smith', '2345678901', 'UnitedHealth', '67890'),
('Alice', 'Jones', 'alice.jones@example.com', '3456789012', '789 Pine Road', strftime('%s', '1990-03-03'), 'f', 'active', '', '', 'AB+', 'Bob Jones', '3456789012', 'Aetna', '11223');