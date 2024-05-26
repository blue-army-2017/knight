import type { Actions } from "./$types";
import * as schema from "$lib/schema";
import { redirect } from "@sveltejs/kit";

export const actions = {
  default: async ({ request, locals: { db } }) => {
    const data = await request.formData();

    const firstName = data.get("first_name");
    const lastName = data.get("last_name");
    const active = data.has("active");

    if (!firstName || !lastName) {
      return { error: "first and last name must be set" };
    }

    await db.insert(schema.members).values({
      id: crypto.randomUUID(),
      firstName: firstName as string,
      lastName: lastName as string,
      active,
    });

    redirect(302, "/members");
  },
} satisfies Actions;
