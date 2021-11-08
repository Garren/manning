<script>
  export let numbers;
  // not reactive
  /* const sum = numbers.reduce((acc, n) => acc + n); */
  $: sum = numbers.reduce((acc, n) => acc + n);
</script>

<div>numbers are {numbers.join(", ")}</div>
<div>sum is {sum}</div>
