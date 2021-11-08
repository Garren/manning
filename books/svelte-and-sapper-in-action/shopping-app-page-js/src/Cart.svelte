<script>
  import { cartStore } from "./stores";
  let total = $cartStore.reduce(
    (acc, item) => acc + item.price * item.quantity,
    0
  );
</script>

<h2>Cart</h2>

{#if $cartStore.length === 0}
  <div>Empty</div>
{:else}
  <table>
    <thead>
      <tr>
        <th>Description</th>
        <th>Quantity</th>
        <th>Price</th>
      </tr>
    </thead>
    <tbody>
      {#each $cartStore as item}
        <tr>
          <td>{item.description}</td>
          <td>{item.quantity}</td>
          <td>{item.price.toFixed(2)}</td>
        </tr>
      {/each}
    </tbody>
    <tfoot>
      <tr>
        <td colspan="2"><label for="">Total</label></td>
        <td>${total.toFixed(2)}</td>
      </tr>
    </tfoot>
  </table>
{/if}
