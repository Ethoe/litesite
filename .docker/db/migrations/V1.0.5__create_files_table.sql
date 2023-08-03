CREATE TABLE IF NOT EXISTS files (
    id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT(6) UNSIGNED NOT NULL,
    filename VARCHAR(255) NOT NULL,
    file_path VARCHAR(255) NOT NULL,
    file_size BIGINT UNSIGNED NOT NULL,
    upload_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- Add additional columns as needed for other metadata
    
    CONSTRAINT fk_user_id
    FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
