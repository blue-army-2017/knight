<script lang="ts">
  import type { ActionData } from "./$types";
  import MemberForm from "$lib/forms/MemberForm.svelte";

  export let form: ActionData;
</script>

{#if form?.error}
  <aside class="alert variant-filled-error">
    <div class="alert-message">
      <h3 class="h3">Error while creating member</h3>
      <p>{form.error}</p>
    </div>
  </aside>
{/if}

<h1 class="my-4">New Member</h1>

<MemberForm />
