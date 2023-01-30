<script>
  import Card from "./card.svelte";
</script>

<Card>
  <p>Smol: https://localhost:3000/r/</p>
  <p>Redirect: {"redirect"}</p>
  <p>Clicked:</p>
  <button class="update">Update</button>
  <button class="delete">Delete</button>
</Card>

<style>
    button {
        color: white;
        font-weight: bold;
        border: none;
        padding: .75rem;
        border-radius: 4px;
    }

    .update {
        background-color: yellowgreen;
    }

    .delete {
        background-color: red;
    }
</style>