create table personalidades(
    id serial primary key,
    nome varchar,
    historia varchar
);

INSERT INTO personalidades (nome, historia) VALUES
('Patricio', 'eu aqui'),
('Patricio 2', 'eu aqui 2');