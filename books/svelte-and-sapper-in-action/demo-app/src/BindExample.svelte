<script>
  const colors = ["red", "orange", "yellow", "green", "blue", "purple"];
  const flavors = ["vanilla", "chocolate", "strawberry"];
  const seasons = ["Spring", "Summer", "Fall", "Winter"];
  let favoriteColor = "";
  let favoriteFlavors = [];
  let favoriteSeason = "";
  let happy = true;
  let name = "";
  let story = "";
</script>

<div class="form">
  <div>
    <label for="name" />
    Name
    <input type="text" name="name" bind:value={name} />
  </div>
  <div>
    <label for="">
      <input type="checkbox" bind:checked={happy} />
      Happy?
    </label>
  </div>
  <div>
    <label for=""> Favorite Flavors </label>
    {#each flavors as flavor}
      <label for="" class="indent">
        <input type="checkbox" value={flavor} bind:group={favoriteFlavors} />
        {flavor}
      </label>
    {/each}
  </div>
  <div>
    <label for="">
      {#each seasons as season}
        <label for="" class="indent">
          <input type="radio" value={season} bind:group={favoriteSeason} />
          {season}
        </label>
      {/each}
    </label>
  </div>
  <div>
    <label for=""
      >Favorite Color
      <select bind:value={favoriteColor}>
        <option />
        {#each colors as color}
          <option>{color}</option>
        {/each}
      </select>
    </label>
  </div>
  <div>
    <label for="">
      Life story
      <textarea bind:value={story} />
    </label>
  </div>
  {#if name}
    <div>
      {name} likes {favoriteColor}, {favoriteSeason}, and is {happy
        ? "happy"
        : "unhappy"}.
    </div>
    <div>{name}'s favorite flavors are {favoriteFlavors}.</div>
    <div>Story: {story}</div>
  {/if}
</div>

<style>
  div {
    margin-bottom: 10px;
  }
  .indent {
    margin-left: 10px;
  }
  input,
  select,
  textarea {
    border: solid lightgray 1px;
    border-radius: 4px;
    margin: 0;
    padding: 4px;
  }
  input[type="checkbox"],
  input[type="radio"] {
    margin: 0 5px 0 0;
  }
  label {
    display: inline-flex;
    align-items: center;
  }
  select,
  textarea {
    margin-left: 5px;
  }
</style>
