import type { PageServerLoad } from "./$types";
import * as schema from "$lib/schema";

export const load: PageServerLoad = async ({ locals: { db } }) => {
  const members = await db
    .select()
    .from(schema.members)
    .orderBy(schema.members.lastName, schema.members.firstName);

  return {
    members,
  };
};
