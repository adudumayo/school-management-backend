CREATE TABLE quiz (
    id INT AUTO_INCREMENT PRIMARY KEY,
    subject VARCHAR(50),
    topic VARCHAR(100),
    question VARCHAR(500),
    due_date DATETIME
);
