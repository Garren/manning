<script>
  let clientH, clientW, offsetH, offsetW;
</script>

<div
  bind:clientHeight={clientH}
  bind:clientWidth={clientW}
  bind:offsetHeight={offsetH}
  bind:offsetWidth={offsetW}
>
  How big am I?
</div>
<ul>
  <li>{clientH}</li>
  <li>{clientW}</li>
  <li>{offsetH}</li>
  <li>{offsetW}</li>
</ul>

<style>
  div {
    background-color: green;
    height: 100px;
    width: 123px;
  }
</style>
