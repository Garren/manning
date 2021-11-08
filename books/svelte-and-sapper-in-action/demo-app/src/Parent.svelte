<script>
  import Child from "./Child.svelte";
  let pValue = 1;
  let triple;
</script>

<h1>Parent</h1>
<Child bind:cValue={pValue} />
<div>pValue = {pValue}</div>

<Child bind:cValue={pValue} bind:triple />
<div>pValue = {pValue}</div>
<div>triple = {triple}</div>
