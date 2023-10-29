-- sample table for transactions
create table transactions
(
    id		int auto_increment,
    name 	varchar(100) not null,
    pt 		varchar(100) not null,
    amount 	int			 not null,
    date 	timestamp 	 not null,
    primary key (id)
) engine = InnoDB;

-- initial transactions data
insert into transactions
values
	(1, 'A', 'X', 15000, '2023/1/1'),
	(2, 'A', 'Y', 10000, '2023/1/1'),
	(3, 'A', 'X', 25000, '2023/2/1'),
	(4, 'B', 'X', 5000, '2023/2/1'),
	(5, 'C', 'X', 50000, '2023/2/1'),
	(6, 'A', 'X', 25000, '2023/3/1'),
	(7, 'A', 'X', 30000, '2023/4/1'),
	(8, 'B', 'X', 5000, '2023/4/1'),
	(9, 'A', 'X', 25000, '2023/5/1');

-- snapshot table to illustrate the snapshot process
create table snapshots
(
    id		int auto_increment,
    date 	timestamp 	 not null,
    amount 	int			 not null,
    primary key (id)
) engine = InnoDB;


-- how snapshot is done in simple operations
truncate table snapshots;

insert into snapshots
select null, date, sum(amount)
from transactions
group by date;
