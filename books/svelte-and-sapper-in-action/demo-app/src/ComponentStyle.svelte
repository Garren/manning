<script>
  let status = 200;
  let message = "this is a message";
  $: error = status >= 400;
</script>

<div>
  <label>
    Status
    <input type="number" bind:value={status} />
  </label>
  <input
    type="range"
    id="range"
    min="100"
    max="500"
    step="100"
    bind:value={status}
  />
</div>

<div class:error>{message}</div>

<!--<div class:error={status >= 400}>{message}</div>-->
<style>
  .error {
    color: red;
    font-weight: bold;
  }
</style>
