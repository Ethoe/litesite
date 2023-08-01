-- Remove the lastname column
ALTER TABLE users DROP COLUMN lastname;
ALTER TABLE users DROP COLUMN age;

-- Change the firstname column to username and make it unique
ALTER TABLE users CHANGE firstname username VARCHAR(30) NOT NULL;
ALTER TABLE users ADD CONSTRAINT unique_username UNIQUE (username);
