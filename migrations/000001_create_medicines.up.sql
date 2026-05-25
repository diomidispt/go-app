CREATE TABLE medicines (
    id      SERIAL          PRIMARY KEY,  -- auto-incrementing unique ID, Postgres generates this automatically
    name    TEXT            NOT NULL,     -- medicine name, cannot be empty
    dosage  TEXT            NOT NULL,     -- e.g. "500mg", "10ml" — free text, cannot be empty
    stock   INT             NOT NULL DEFAULT 0,  -- how many units in stock, defaults to 0 if not provided
    price   NUMERIC(10,2)   NOT NULL      -- price with up to 10 digits and 2 decimal places (e.g. 12.99)
);
