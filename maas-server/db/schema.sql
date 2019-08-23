CREATE TABLE `items` (
	`ID` INTEGER PRIMARY KEY AUTOINCREMENT,
	`name` VARCHAR(256) NOT NULL,
	`cost` INTEGER
);
CREATE TABLE `items_stats` (
	`itemID` INTEGER PRIMARY KEY,
	`consumed` INTEGER NOT NULL
);
CREATE TABLE `user_items_stats` (
  `userID`INTEGER,
  `itemID` INTEGER not null,
  `consumed` INTEGER not null,
  constraint user_items_stats_pk
  primary key (userID, itemID)
);
CREATE TABLE `users` (
	`ID` INTEGER PRIMARY KEY AUTOINCREMENT,
	`username` VARCHAR(256) NOT NULL,
	`password` VARCHAR(256) NOT NULL,
	`credits` INTEGER NOT NULL,
	`admin` INTEGER NOT NULL
);
CREATE TABLE `item_inventory` (
	`itemID` INTEGER PRIMARY KEY,
	`stock` INTEGER
);
