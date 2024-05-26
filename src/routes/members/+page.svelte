<script lang="ts">
  import type { PageData } from "./$types";
  import { goto } from "$app/navigation";

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
      </tr>
    </thead>
    <tbody>
      {#each data.members as member}
        <tr class="hover:cursor-pointer" on:click={() => goto(`/members/${member.id}`)}>
          <td>{member.lastName}</td>
          <td>{member.firstName}</td>
          {#if member.active}
            <td>
              <span class="badge variant-filled-success">Active</span>
            </td>
          {:else}
            <td>
              <span class="badge variant-filled-error">Inactive</span>
            </td>
          {/if}
        </tr>
      {/each}
    </tbody>
  </table>
</div>

<a href="/members/new" class="btn variant-filled-primary m-4">New Member</a>
