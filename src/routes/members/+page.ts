import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ url, fetch }) => {
  const response = await fetch(url.pathname);

  return {
    members: await response.json(),
  };
};
