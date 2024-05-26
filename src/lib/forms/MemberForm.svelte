<script lang="ts">
  export let firstName = "";
  export let lastName = "";
  export let active = true;

  $: isValid = !!firstName.trim() && !!lastName.trim();
</script>

<form method="post" class="px-4 space-y-4">
  <label class="label">
    <span>First Name</span>
    <input class="input" name="first_name" type="text" placeholder="Max" bind:value={firstName} />
  </label>

  <label class="label">
    <span>Last Name</span>
    <input
      class="input"
      name="last_name"
      type="text"
      placeholder="Mustermann"
      bind:value={lastName}
    />
  </label>

  <label class="flex items-center space-x-2">
    <input class="checkbox" name="active" type="checkbox" bind:checked={active} />
    <p>Active</p>
  </label>

  <section class="flex justify-evenly items-center">
    <button type="submit" class="btn variant-filled-primary" disabled={!isValid}>Save</button>
    <a href="/members" class="btn variant-filled-secondary">Back</a>
  </section>
</form>
