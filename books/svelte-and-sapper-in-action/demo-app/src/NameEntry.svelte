<script>
  import { onMountFocus, onMountLog } from "./helper.js";
  export let name; // prop
  onMountLog("NameEntry");
  onMountFocus();
</script>

<label>
  Name
  <input bind:value={name} />
</label>
