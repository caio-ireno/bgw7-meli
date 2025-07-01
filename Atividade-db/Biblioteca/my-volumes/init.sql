DROP DATABASE IF EXISTS biblioteca;
CREATE DATABASE biblioteca CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE biblioteca;

DROP TABLE IF EXISTS `livro`;

CREATE TABLE livro(

`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`titulo` varchar(100) NOT NULL,
`editorial` varchar(100) NOT NULL,
`area` varchar(100) NOT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


DROP TABLE IF EXISTS `autor`;
CREATE TABLE autor(
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`nome` varchar(100) NOT NULL,
`nacionalidade` varchar(100) NOT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `estudante`;
CREATE TABLE estudante(
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`nome` varchar(100) NOT NULL,
`sobrenome` varchar(100) NOT NULL,
`endereco` varchar(100) NOT NULL,
`carreira` varchar(100) NOT NULL,
`idade` int(10) NOT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


DROP TABLE IF EXISTS `livro_autor`;
CREATE TABLE livro_autor(
`idLivro` int(10) unsigned NOT NULL,
`idAutor` int(10) unsigned NOT NULL,
PRIMARY KEY (`idLivro`, `idAutor`),
FOREIGN KEY (`idLivro`) REFERENCES `livro`(`id`),
FOREIGN KEY (`idAutor`) REFERENCES `autor`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=    utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `emprestimo`;
CREATE TABLE emprestimo(
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`idEstudante` int(10) unsigned NOT NULL,
`idLivro` int(10) unsigned NOT NULL,
`dataEmprestimo` date NOT NULL,
`dataDevolucao` date NOT NULL,
`devolvido` boolean NOT NULL,
PRIMARY KEY (`id`),
FOREIGN KEY (`idLivro`) REFERENCES `livro`(`id`),
FOREIGN KEY (`idEstudante`) REFERENCES `estudante`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=    utf8mb4_unicode_ci;


INSERT INTO autor (nome, nacionalidade) VALUES
  ('Machado de Assis', 'Brasileiro'),
  ('José Saramago', 'Português'),
  ('Clarice Lispector', 'Brasileira'),
  ('George Orwell', 'Britânico');

INSERT INTO livro (titulo, editorial, area) VALUES
  ('Dom Casmurro', 'Editora Globo', 'Ficção'),
  ('Ensaio sobre a Cegueira', 'Caminho', 'Romance'),
  ('A Hora da Estrela', 'Rocco', 'Ficção'),
  ('1984', 'Companhia das Letras', 'Distopia'),
  ('Memórias Póstumas de Brás Cubas', 'Editora Globo', 'Ficção');

INSERT INTO estudante (nome, sobrenome, endereco, carreira, idade) VALUES
  ('Ana', 'Silva', 'Rua das Flores, 100', 'Engenharia', 21),
  ('Bruno', 'Souza', 'Avenida Central, 45', 'Direito', 22),
  ('Carla', 'Pereira', 'Travessa Alegre, 12', 'Medicina', 20);

INSERT INTO livro_autor (idLivro, idAutor) VALUES
  (1, 1), 
  (2, 2), 
  (3, 3), 
  (4, 4), 
  (5, 1); 

INSERT INTO emprestimo (idEstudante, idLivro, dataEmprestimo, dataDevolucao, devolvido) VALUES
  (1, 1, '2023-04-01', '2023-04-15', TRUE), 
  (1, 2, '2023-05-10', '2023-05-24', FALSE), 
  (2, 3, '2023-03-20', '2023-04-03', TRUE),   
  (3, 4, '2023-02-25', '2023-03-11', TRUE),   
  (2, 5, '2023-06-01', '2023-06-15', FALSE); 