create table games (
	game_id int primary key auto_increment,
	title varchar(255) not null,
	genre varchar(100) not null,
	price float not null,
	stock int not null
);

create table branches (
	branch_id int primary key auto_increment,
	name varchar(255) not null,
	location varchar(255) not null
);

create table sales (
	sale_id int primary key auto_increment,
	game_id int,
	branch_id int,
	sale_date DATE NOT NULL,
	quantity INT NOT NULL,
	FOREIGN KEY (game_id) REFERENCES games(game_id),
	FOREIGN KEY (branch_id) REFERENCES branches(branch_id)
);


INSERT INTO games (title, genre, price, stock) VALUES
('Final Fantasy', 'RPG', 50.99, 100),
('FIFA 23', 'Sports', 49.99, 120),
('Doom Eternal', 'FPS', 49.99, 80);

INSERT INTO branches (name, location) VALUES
('Downtown Branch', '123 Downtown St'),
('Uptown Branch', '456 Uptown St');

INSERT INTO sales (game_id, branch_id, sale_date, quantity) VALUES
(1, 1, '2023-08-26', 2),
(2, 2, '2023-08-25', 3);
