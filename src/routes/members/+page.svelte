<script lang="ts">
  import type { PageData } from "./$types";

  export let data: PageData;
</script>

<h1 class="m-4">Members</h1>

<div class="table-container px-4">
  <table class="table table-hover">
    <thead>
      <tr>
        <th>Last Name</th>
        <th>First Name</th>
        <th>Active</th>
        <th>Actions</th>
      </tr>
    </thead>
    <tbody>
      {#each data.members as member}
        <tr>
          <td>{member.lastName}</td>
          <td>{member.firstName}</td>
          <td>{member.active ? "yes" : "no"}</td>
          <td class="table-cell-fit">
            <div class="btn-group">
              <button class="variant-filled-secondary">Edit</button>
              <button class="variant-filled-error">Delete</button>
            </div>
          </td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>

<button class="btn variant-filled-primary m-4">New Member</button>
