<script>
  import items from "./items";
  import { cartStore } from "./stores";

  function changeQuantity(event, item) {
    const newQuantity = Number(event.target.value);
    cartStore.update((items) => {
      if (newQuantity && !item.quantity) {
        // new quantity is not zero but the old quantity is...
        items.push(item);
      } else if (newQuantity === 0 && item.quantity) {
        // new quantity is zero  but the old quantity is not
        const { description } = item;
        items = items.filter((i) => i.description !== description);
      }
      item.quantity = newQuantity;
      return items; // have to return the modified state
    });
  }
</script>

<h1>Shop</h1>

<table>
  <thead>
    <tr>
      <th>Description</th>
      <th>Price</th>
      <th>Quantity</th>
    </tr>
  </thead>
  <tbody>
    {#each items as item}
      <tr>
        <td>{item.description}</td>
        <td>${item.price.toFixed(2)}</td>
        <td>
          <input
            type="number"
            min="0"
            on:input={(e) => changeQuantity(e, item)}
            value={item.quantity}
          />
        </td>
      </tr>
    {/each}
  </tbody>
</table>

<style>
  input {
    width: 60px;
  }
</style>
