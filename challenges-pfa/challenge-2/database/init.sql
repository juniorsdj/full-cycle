CREATE DATABASE pfa;
USE pfa;

CREATE TABLE module(
    codigo int(4) AUTO_INCREMENT,
    nome varchar(30) NOT NULL,
    PRIMARY KEY (codigo)
);

INSERT INTO module( nome) VALUES ("Docker") ;
INSERT INTO module( nome) VALUES ("Fundamentos de Arquitetura") ;
INSERT INTO module( nome) VALUES ("Comunicação") ;