<script>
  import dialogPolyfill from "dialog-polyfill";
  import { createEventDispatcher, onMount } from "svelte";

  export let canClose = true;
  export let className = "fixed";
  export let dialog = null;
  export let icon = undefined;
  export let title;

  const dispatch = createEventDispatcher();

  $: classNames = "dialog" + (className ? " " + className : "");

  onMount(() => {
    // lifecycle functions can't be called conditinally, so we
    // have to move our conditinal check into onMount.
    if (dialogPolyfill) dialogPolyfill.registerDialog(dialog);
  });

  function close() {
    dispatch("close");
    dialog.close();
  }
</script>

<dialog bind:this={dialog} class={classNames} open>
  <header>
    {#if icon} {icon} {/if}
    <div class="title">{title}</div>
    {#if canClose}
      <button class="close-btn" on:click={close}>&#x2716;</button>
    {/if}
  </header>
  <main>
    <slot />
  </main>
</dialog>

<style>
  .body {
    padding: 10px;
  }

  .close-btn {
    background-color: transparent;
    border: none;
    color: white;
    cursor: pointer;
    font-size: 24px;
    outline: none;
    margin: 0;
    padding: 0;
  }

  dialog {
    position: fixed;
    top: 50%;
    transform: translate(0, -50%);
    border: none;
    box-shadow: 0 0 10px darkgray;
    padding: 0;
  }

  header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: cornflowerblue;
    box-sizing: border-box;
    color: white;
    font-weight: bold;
    padding: 10px;
    width: 100%;
  }

  main {
    padding: 10px;
  }

  .title {
    flex-grow: 1;
    font-size: 18px;
    margin-right: 10px;
  }

  dialog::backdrop,
  :global(dialog + .backdrop) {
    background: rgba(0, 0, 0, 0.4);
  }
</style>
