CREATE DATABASE 'book_store';

CREATE USER 'pnguyen'@'%' IDENTIFIED BY 'deptraithimoiconhieuduaiu';
GRANT ALL PRIVILEGES ON * . * TO 'pnguyen'@'%';
FLUSH PRIVILEGES;