<script lang="ts">
  export let data = {
    firstName: "",
    lastName: "",
    active: true,
  };
  export let canDelete = false;

  $: isValid = !!data.firstName.trim() && !!data.lastName.trim();
</script>

<form method="post" action="?/save" class="px-4 space-y-4">
  <label class="label">
    <span>First Name</span>
    <input
      class="input"
      name="first_name"
      type="text"
      placeholder="Max"
      bind:value={data.firstName}
    />
  </label>

  <label class="label">
    <span>Last Name</span>
    <input
      class="input"
      name="last_name"
      type="text"
      placeholder="Mustermann"
      bind:value={data.lastName}
    />
  </label>

  <label class="flex items-center space-x-2">
    <input class="checkbox" name="active" type="checkbox" bind:checked={data.active} />
    <p>Active</p>
  </label>

  <section class="flex justify-evenly items-center">
    <button type="submit" class="btn variant-filled-primary" disabled={!isValid}>Save</button>
    {#if canDelete}
      <button type="submit" formaction="?/delete" class="btn variant-filled-error">Delete</button>
    {/if}
    <a href="/members" class="btn variant-filled-secondary">Back</a>
  </section>
</form>
