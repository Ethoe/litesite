CREATE TABLE IF NOT EXISTS pages (
    id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    pagename VARCHAR(50),
    reg_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);