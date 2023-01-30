<script>
  import Item from "./item.svelte";
  import { onMount } from "svelte";

  let smols = [];

  onMount(async () => {
    const res = await fetch("http://localhost:3000/smol");
    smols = await res.json();
  });
</script>

{#each smols as smol}
  <Item smol={smol}/>
{/each}
