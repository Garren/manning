<script>
  import { onMountLog } from "./helper.js";
  import AgeEntry from "./AgeEntry.svelte";
  import NameEntry from "./NameEntry.svelte";

  let age = 0;
  let enterAge = false;
  let name = "";
  onMountLog("HelperDemo");
</script>

{#if enterAge}
  <AgeEntry bind:age />
{:else}
  <NameEntry bind:name />
{/if}

<label>
  Enter Age?
  <input type="checkbox" bind:checked={enterAge} />
</label>

<div>{name} is {age} years old.</div>
