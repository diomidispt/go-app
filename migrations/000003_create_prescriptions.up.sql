CREATE TABLE prescriptions (
    id          SERIAL  PRIMARY KEY,                             -- auto-incrementing unique ID
    patient_id  INT     NOT NULL REFERENCES patients(id),       -- links to the patients table — cannot create a prescription for a non-existent patient
    medicine_id INT     NOT NULL REFERENCES medicines(id),      -- links to the medicines table — cannot prescribe a non-existent medicine
    date        DATE    NOT NULL DEFAULT CURRENT_DATE,          -- date of prescription, defaults to today
    instructions TEXT   NOT NULL                                -- e.g. "take one tablet twice a day with food"
);
