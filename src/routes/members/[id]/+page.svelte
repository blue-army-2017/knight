<script lang="ts">
  import MemberForm from "$lib/forms/MemberForm.svelte";
  import type { ActionData, PageData } from "./$types";

  export let data: PageData;
  export let form: ActionData;
</script>

{#if form?.error}
  <aside class="alert variant-filled-error">
    <div class="alert-message">
      <h3 class="h3">Error while editing member</h3>
      <p>{form.error}</p>
    </div>
  </aside>
{/if}

<h1 class="m-4">Edit Member</h1>

<MemberForm data={data.member} canDelete />
