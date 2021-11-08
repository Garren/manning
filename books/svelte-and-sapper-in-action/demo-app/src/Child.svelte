<script>
  export let cValue;
  export let triple;
  const double = () => (cValue *= 2);
  $: triple = cValue * 3;
</script>

<h2>Child</h2>
<button on:click={double}>Double</button>
<div>cValue={cValue}</div>
