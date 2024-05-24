import { json, error } from "@sveltejs/kit";
import type { RequestHandler } from "@sveltejs/kit";
import * as schema from "$lib/schema";

export const GET: RequestHandler = async ({ locals: { db } }) => {
  try {
    const members = await db
      .select()
      .from(schema.members)
      .orderBy(schema.members.lastName, schema.members.firstName);

    return json(members);
  } catch (e) {
    error(500);
  }
};
