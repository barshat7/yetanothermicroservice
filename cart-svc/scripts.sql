create table cart (
id int not null auto_increment,
user_id varchar(40),
active boolean,
created_date datetime,
primary key (id)
);
--------------------------------------
create table cart_items (
	id int not null auto_increment,
    cart_id int,
    item_id int,
    item_name varchar(255),
    item_price double,
    primary key (id),
    foreign key (cart_id) references cart(id)
);