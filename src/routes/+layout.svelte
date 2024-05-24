<script lang="ts">
  import "../app.postcss";
  import { AppBar, LightSwitch } from "@skeletonlabs/skeleton";
  import logo from "$lib/assets/logo.png";

  // Floating UI for Popups
  import { computePosition, autoUpdate, flip, shift, offset, arrow } from "@floating-ui/dom";
  import { storePopup } from "@skeletonlabs/skeleton";
  storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });
</script>

<AppBar class="w-full">
  <svelte:fragment slot="lead">
    <img src={logo} alt="Blue Army Logo" class="h-8" />
  </svelte:fragment>
  <a href="/" class="font-bold no-underline">Knight</a>
  <svelte:fragment slot="trail">
    <LightSwitch />
  </svelte:fragment>
</AppBar>

<main class="container h-full flex flex-col justify-center items-center">
  <slot />
</main>
