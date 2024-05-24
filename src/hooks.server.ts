import { Database } from "bun:sqlite";
import { drizzle } from "drizzle-orm/bun-sqlite";
import { migrate } from "drizzle-orm/bun-sqlite/migrator";
import * as schema from "$lib/schema";
import type { Handle } from "@sveltejs/kit";

const sqlite = new Database(Bun.env.DB_PATH || "knight.db");
const db = drizzle(sqlite, { schema });

await migrate(db, { migrationsFolder: "./drizzle" });

export const handle: Handle = async ({ event, resolve }) => {
  event.locals.db = db;

  return await resolve(event);
};
