CREATE TABLE `present_members` (
	`member_id` text NOT NULL,
	`game_id` text NOT NULL,
	PRIMARY KEY(`game_id`, `member_id`),
	FOREIGN KEY (`member_id`) REFERENCES `members`(`id`) ON UPDATE no action ON DELETE no action,
	FOREIGN KEY (`game_id`) REFERENCES `season_games`(`id`) ON UPDATE no action ON DELETE no action
);
--> statement-breakpoint
CREATE TABLE `season_games` (
	`id` text PRIMARY KEY NOT NULL,
	`opponent` text NOT NULL,
	`home` integer NOT NULL,
	`mode` text NOT NULL,
	`date` text NOT NULL,
	`season_id` text NOT NULL,
	FOREIGN KEY (`season_id`) REFERENCES `seasons`(`id`) ON UPDATE no action ON DELETE no action
);
--> statement-breakpoint
CREATE TABLE `seasons` (
	`id` text PRIMARY KEY NOT NULL,
	`name` text NOT NULL,
	`created` text NOT NULL
);
