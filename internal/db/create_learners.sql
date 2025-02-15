DROP TABLE IF EXISTS learner;

CREATE TABLE learner(
    id      INT AUTO_INCREMENT NOT NULL,
    name    VARCHAR(128) NOT NULL,
    class   INT NOT NULL,
    average DECIMAL(5,2) NOT NULL,
    PRIMARY KEY (`id`)
);
