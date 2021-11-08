<script>
  import { afterUpdate, beforeUpdate } from "svelte";
  let input, name, selectionEnd, selectionStart;

  beforeUpdate(() => {
    if (input) ({ selectionStart, selectionEnd } = input);
  });

  afterUpdate(() => {
    input.setSelectionRange(selectionStart, selectionEnd);
    input.focus();
  });
</script>

<input bind:this={input} bind:value={name} />
<!-- hitting the UPPER button modifies the name state variable causing
svelte to schedule it for re-rendering. Rendering will lose the selected
range in the control, so before we update the component we capture it. When 
the component has been updated by svelte, then we reset it.
-->
<button on:click={() => (name = name.toUpperCase())}>UPPER</button>
