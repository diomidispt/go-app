ALTER TABLE medicines ADD COLUMN deleted_at TIMESTAMP; -- null means active, a timestamp means soft deleted
