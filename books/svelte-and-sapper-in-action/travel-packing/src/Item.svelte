<script>
  import { blurOnKey } from "./util";
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();
  export let item;
  let editing = false;

  $: if (editing) console.log("editing item", item.name);

  export let categoryId; // so each item knows its current category
  export let dnd; // a propr that will recieve an object that has dnd methods
</script>

<li>
  <input type="checkbox" bind:checked={item.packed} />
  {#if editing}
    <!-- svelte-ignore a11y-autofocus -->
    <input
      type="text"
      autofocus
      bind:value={item.name}
      on:blur={() => (editing = false)}
      on:keydown={blurOnKey}
    />
  {:else}
    <span
      class="packed-{item.packed}"
      on:click={() => (editing = true)}
      draggable={true}
      on:dragstart={(event) => dnd.drag(event, categoryId, item.id)}
    >
      {item.name}
    </span>
  {/if}
  <button data-testid="delete" class="icon" on:click={() => dispatch("delete")}
    >&#x1F5D1;</button
  >
</li>

<style>
  button {
    background-color: transparent;
    border: none;
  }
  input[type="checkbox"] {
    --size: 1.5rem;
    height: var(--size);
    width: var(--size);
  }
  input[type="text"] {
    border: solid lightgray 1px;
  }
  li {
    display: flex;
    align-items: center;
  }
  .packed-true {
    color: gray;
    text-decoration: line-through;
  }
  span {
    margin: 0 0.5rem;
  }
</style>
