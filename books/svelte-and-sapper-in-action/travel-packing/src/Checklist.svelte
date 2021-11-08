<script>
  import { createEventDispatcher } from "svelte";
  import Category from "./Category.svelte";
  import { getGuid, sortOnName } from "./util";
  import Dialog from "./Dialog.svelte";
  import { flip } from "svelte/animate";

  const dispatch = createEventDispatcher();
  const options = { duration: 700 };

  let dialog = null;

  let categoryArray = [];
  let categories = {};
  let categoryName;
  let message = "";
  let show = "all";

  let dragAndDrop = {
    drag(event, categoryId, itemId) {
      const data = { categoryId, itemId };
      event.dataTransfer.setData("text/plain", JSON.stringify(data));
      event.dataTransfer.dropEffect = "move";
    },
    drop(event, categoryId) {
      const json = event.dataTransfer.getData("text/plain");
      const data = JSON.parse(json);

      const category = categories[data.categoryId];
      const item = category.items[data.itemId];
      delete category.items[data.itemId];

      categories[categoryId].items[data.itemId] = item;

      categories = categories;
    },
  };

  $: categoryArray = sortOnName(Object.values(categories));

  function deleteCategory(category) {
    if (Object.values(category.items).length) {
      message = "This category is not empty";
      dialog.showModal();
      return;
    }
    delete categories[category.id];
    categories = categories;
  }

  function addCategory() {
    const duplicate = Object.values(categories).some(
      (cat) => cat.name === categoryName
    );
    if (duplicate) {
      message = `The category "${categoryName}" already exists.`;
      dialog.showModal();
      return;
    }
    const id = getGuid();
    categories[id] = { id, name: categoryName, items: {} };
    categories = categories; // triggers an update
    categoryName = ""; // clears the input
  }
  function clearAllChecks() {
    for (const category of Object.values(categories)) {
      for (const item of Object.values(category.items)) {
        item.packed = false;
      }
    }
    categories = categories; // triggers an update
  }

  restore();

  $: if (categories) persist(); // when categories changes...

  function persist() {
    localStorage.setItem("travel-packing", JSON.stringify(categories));
  }
  function restore() {
    const text = localStorage.getItem("travel-packing");
    if (text && text !== "{}") {
      categories = JSON.parse(text);
    }
  }
</script>

<section>
  <header>
    <form on:submit|preventDefault={addCategory}>
      <label>
        New Category
        <input
          data-testid="category-name-input"
          required
          bind:value={categoryName}
        />
      </label>
      <button disabled={!categoryName}>Add Category</button>
      <button class="logout-btn" on:click={() => dispatch("logout")}
        >Log Out</button
      >
    </form>
    <p>
      Suggested categories include Backpack, Clothes,
      <br />
      Last Minute, Medicines, Running Gear, and Toiletries
    </p>
    <div class="radios">
      <label for=""> Show </label>
      <label>
        <input type="radio" value="all" name="show" bind:group={show} />
        All
      </label>
      <label>
        <input type="radio" value="packed" name="show" bind:group={show} />
        Packed
      </label>
      <label>
        <input type="radio" value="unpacked" name="show" bind:group={show} />
        Unpacked
      </label>
      <button class="clear" on:click={clearAllChecks}>Clear All Checks</button>
    </div>
  </header>
  <div class="categories">
    {#each categoryArray as category (category.id)}
      <div class="wrapper" animate:flip={options}>
        <Category
          bind:category
          {categories}
          {show}
          on:persist={persist}
          on:delete={() => deleteCategory(category)}
          dnd={dragAndDrop}
        />
      </div>
    {/each}
  </div>
</section>
<Dialog title="Checklist" bind:dialog>
  <div>{message}</div>
</Dialog>

<style>
  .animate {
    display: inline-block;
  }
  .wrapper {
    display: inline;
  }
  .categories {
    display: inline-flex;
    flex-wrap: wrap;
    justify-content: center;
  }
  .clear {
    margin-left: 2rem;
  }
  input[type="radio"] {
    --size: 1.5rem;
    height: var(--size);
    width: var(--size);
  }
  .logout-btn {
    position: absolute;
    right: 1rem;
    top: 1rem;
  }
  .radios {
    display: flex;
    align-items: center;
  }
  .radios > label:not(:first-of-type) {
    display: inline-flex;
    align-items: center;
    margin-left: 1rem;
  }
  /* .radios > label > input { */
  /*   margin-bottom: -3px; */
  /*   margin-right: 5px; */
  /* } */
  section {
    display: flex;
    flex-direction: column;
    align-items: center;
    font-size: 1.5rem;
    margin-top: 1rem;
  }
</style>
