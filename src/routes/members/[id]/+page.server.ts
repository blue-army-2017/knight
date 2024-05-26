import type { Actions, PageServerLoad } from "./$types";
import * as schema from "$lib/schema";
import { error, fail, redirect } from "@sveltejs/kit";
import { eq } from "drizzle-orm";

export const load: PageServerLoad = async ({ params, locals: { db } }) => {
  const result = await db.select().from(schema.members).where(eq(schema.members.id, params.id));
  if (result.length < 1) {
    return error(404, "member not found");
  }

  return {
    member: result[0],
  };
};

export const actions: Actions = {
  save: async ({ request, params, locals: { db } }) => {
    const data = await request.formData();

    const firstName = data.get("first_name");
    const lastName = data.get("last_name");
    const active = data.has("active");

    if (!firstName || !lastName) {
      return fail(400, { error: "first and last name must be set" });
    }

    await db
      .update(schema.members)
      .set({
        firstName: firstName as string,
        lastName: lastName as string,
        active,
      })
      .where(eq(schema.members.id, params.id));

    redirect(302, "/members");
  },
  delete: async ({ params, locals: { db } }) => {
    await db.delete(schema.members).where(eq(schema.members.id, params.id));

    redirect(302, "/members");
  },
};
