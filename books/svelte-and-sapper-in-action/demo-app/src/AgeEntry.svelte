<script>
  import { onMountFocus, onMountLog } from "./helper.js";
  export let age; // prop
  onMountLog("AgeEntry");
  onMountFocus();
</script>

<label>
  Age
  <input bind:value={age} type="number" min="0" />
</label>
