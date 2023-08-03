-- Add a UNIQUE constraint to the file_path column
ALTER TABLE files ADD CONSTRAINT unique_file_path UNIQUE (file_path);

-- Add the url column to the files table
ALTER TABLE files ADD url VARCHAR(255) NOT NULL;
