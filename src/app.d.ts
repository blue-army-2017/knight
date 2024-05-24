// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
// and what to do when importing types
import type { BunSQLiteDatabase } from "drizzle-orm/bun-sqlite";
import type * as schema from "$lib/schema";

declare global {
  declare namespace App {
    interface Locals {
      db: BunSQLiteDatabase<schema>;
    }
    // interface PageData {}
    // interface Error {}
    // interface Platform {}
  }
}
