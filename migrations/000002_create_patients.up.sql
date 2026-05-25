CREATE TABLE patients (
    id          SERIAL  PRIMARY KEY,  -- auto-incrementing unique ID
    name        TEXT    NOT NULL,     -- full name of the patient
    date_of_birth DATE  NOT NULL,     -- stored as a date, not text, so we can calculate age and sort
    phone       TEXT    NOT NULL,     -- contact phone number
    email       TEXT    NOT NULL      -- contact email address
);
