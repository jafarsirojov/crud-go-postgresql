INSERT INTO burgers (name, price)
VALUES ('Big Mac', 16000),
       ('Chicken Mac', 12000),
       ('Chicken Mac55', 12000);


select * from burgers;

UPDATE burgers SET removed=false WHERE id=1;

