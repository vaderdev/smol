<script>
  import Card from "./card.svelte";
  import { Modals, closeModal, openModal } from "svelte-modals"
  export let smol
  async function update(data) {
    const json = {
        redirect: data.redirect,
        smol: data.smol,
        random: data.random,
        id: smol.id
    }
  }
</script>

<Card>
  <p>Smol: https://localhost:3000/r/{smol.smol}</p>
  <p>Redirect: {smol.redirect}</p>
  <p>Clicked: {smol.clicked}</p>
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