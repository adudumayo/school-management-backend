DROP TABLE IF EXISTS learner;

CREATE TABLE learner(
    id      INT AUTO_INCREMENT NOT NULL,
    name    VARCHAR(128) NOT NULL,
    class   INT NOT NULL,
    average DECIMAL(5,2) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO learner (name, class, average) VALUES
('Sipho Ndlovu', 10, 75.50),
('Ayanda Dlamini', 11, 82.30),
('Lebo Mokoena', 9, 90.10),
('Thabo Khumalo', 12, 88.75),
('Zanele Nkosi', 10, 79.20);

