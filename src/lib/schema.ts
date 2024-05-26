import { sqliteTable, text, integer, primaryKey } from "drizzle-orm/sqlite-core";

export const members = sqliteTable("members", {
  id: text("id").primaryKey(),
  firstName: text("first_name").notNull(),
  lastName: text("last_name").notNull(),
  active: integer("active", { mode: "boolean" }).notNull(),
});

export const seasons = sqliteTable("seasons", {
  id: text("id").primaryKey(),
  name: text("name").notNull(),
  created: text("created").notNull(),
});

export const seasonGames = sqliteTable("season_games", {
  id: text("id").primaryKey(),
  opponent: text("opponent").notNull(),
  home: integer("home", { mode: "boolean" }).notNull(),
  mode: text("mode", { enum: ["regular", "playoffs"] }).notNull(),
  date: text("date").notNull(),
  seasonId: text("season_id")
    .notNull()
    .references(() => seasons.id),
});

export const presentMembers = sqliteTable(
  "present_members",
  {
    memberId: text("member_id")
      .notNull()
      .references(() => members.id),
    gameId: text("game_id")
      .notNull()
      .references(() => seasonGames.id),
  },
  (table) => ({
    pk: primaryKey({ columns: [table.memberId, table.gameId] }),
  }),
);
