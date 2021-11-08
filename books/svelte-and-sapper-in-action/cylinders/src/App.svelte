<script>
  let diameter = 1;
  let height = 1;

  $: volume = area * height;
  $: area = Math.PI * radius ** 2;
  $: radius = diameter / 2;
</script>

<h1>Cylinder Calculations</h1>
<label for="diameter">
  Diameter:
  <input id="diameter" type="number" bind:value={diameter} />
</label>
<label for="height">
  Height:
  <input id="height" type="number" bind:value={height} />
</label>

Radius: {radius}
Area: {area.toFixed(2)}
Volume: {volume.toFixed(2)}

<style>
  input {
    width: 50px;
  }
</style>

