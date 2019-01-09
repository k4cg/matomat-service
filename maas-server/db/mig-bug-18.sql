create table user_items_stats_dg_tmp
					(
					userID INTEGER,
					itemID INTEGER not null,
					consumed INTEGER not null,
					constraint user_items_stats_pk
					primary key (userID, itemID)
					);
insert into user_items_stats_dg_tmp(userID, itemID, consumed) select userID, itemID, consumed from user_items_stats;
drop table user_items_stats;
alter table user_items_stats_dg_tmp rename to user_items_stats;
