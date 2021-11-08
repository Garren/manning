<script>
  let breed = "Beagle";

  async function getDogs() {
    const url =
      "https://dog.ceo/api/breed/" + `${breed.toLowerCase()}/images/random/1`;
    const res = await fetch(url);
    if (!res.ok || res.status === 404) return [];
    const json = await res.json();
    return json.message;
  }

  let dogsPromise = getDogs();
</script>

<main>
  <label>
    Breed
    <input bind:value={breed} />
  </label>
  <button on:click={() => (dogsPromise = getDogs())}>Get Image</button>
  {#await dogsPromise}
    <div>Wiating for dogs...</div>
  {:then imageUrls}
    {#each imageUrls as url}
      <div><img alt="dog" src={url} /></div>
    {:else}
      <div>not found</div>
    {/each}
  {:catch error}
    <div>Error: {error.message}</div>
  {/await}
</main>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>

