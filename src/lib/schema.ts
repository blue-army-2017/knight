import { sqliteTable, text, integer } from "drizzle-orm/sqlite-core";

export const members = sqliteTable("members", {
  id: text("id").primaryKey(),
  firstName: text("first_name"),
  lastName: text("last_name"),
  active: integer("active", { mode: "boolean" }),
});
