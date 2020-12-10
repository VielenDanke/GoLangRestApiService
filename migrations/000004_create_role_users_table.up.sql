ALTER TABLE users ADD COLUMN authority int;
UPDATE users SET authority=1;
ALTER TABLE users ALTER COLUMN authority SET NOT NULL;