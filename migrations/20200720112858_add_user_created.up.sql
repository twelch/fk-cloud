ALTER TABLE fieldkit.user ADD COLUMN created_at TIMESTAMP DEFAULT NOW();
ALTER TABLE fieldkit.user ADD COLUMN updated_at TIMESTAMP DEFAULT NOW();
