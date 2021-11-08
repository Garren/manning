<script>
  import { afterUpdate } from "svelte";
  let input; // bound to the input DOM element
  let item = "";
  let items = [];
  let list; // bound to the div.list DOM element

  afterUpdate(() => list.scrollTo(0, list.scrollHeight));

  function addItem() {
    // has access to a state variable
    items.push(item);
    items = items; // trigger update
    item = "";
    input.focus(); // because we have a reference to the DOM element
  }
</script>

<p>Tell me what you want for your birthday.</p>
<form on:submit|preventDefault>
  <input bind:this={input} bind:value={item} />
  <button on:click={addItem}>Add</button>
</form>

<div class="list" bind:this={list}>
  {#each items as item}
    <div>{item}</div>
  {/each}
</div>

<style>
  .list {
    border: solid gray 2px;
    height: 52px;
    overflow-y: scroll;
    padding: 5px;
  }
</style>
