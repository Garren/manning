<script>
  import { onMount } from "svelte";
  export let text;
  const colors = ["red", "orange", "yellow", "green", "blue", "purple"];
  let colorIndex = 0;

  onMount(() => {
    let token = setInterval(() => {
      colorIndex = (colorIndex + 1) % colors.length;
    }, 500);
    // return a function and it will be called when the component is destroyed.
    return () => {
      console.log("ColorCycle destroyed");
      clearInterval(token);
    };
  });

  /* onDestroy(() => { */
  /*   console.log("ColorCycle destroyed"); */
  /*   clearInterval(token); */
  /* }); */
</script>

<h2 style="color: {colors[colorIndex]}">{text}</h2>
