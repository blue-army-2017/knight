import { Database } from "bun:sqlite";
import { drizzle } from "drizzle-orm/bun-sqlite";
import { migrate } from "drizzle-orm/bun-sqlite/migrator";
import * as schema from "./schema";

const sqlite = new Database(Bun.env.DB_PATH || "knight.db");
const db = drizzle(sqlite, { schema });

await migrate(db, { migrationsFolder: "./drizzle" });
