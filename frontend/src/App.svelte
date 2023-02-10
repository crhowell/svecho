<script lang="ts">
  import { onMount } from 'svelte';
  import svelteLogo from './assets/svelte.svg'
  import Counter from './lib/Counter.svelte'

  let data;
  $: {
    console.log("data\n", data)
  }
  onMount(() => {
    // PROD: fetch('/api')
    // DEV: caddy proxies to echo -> fetch('http://localhost/api')
    // DEV: direct to echo -> fetch('http://localhost:8000/api')
    // fetch('http://localhost/api')
    fetch('/api')
      .then(resp => data = resp.json())
      .then(json => data = json)
      .catch(err => console.log(err))
  })
</script>

<main>
  <div>
    <a href="https://vitejs.dev" target="_blank" rel="noreferrer"> 
      <img src="/vite.svg" class="logo" alt="Vite Logo" />
    </a>
    <a href="https://svelte.dev" target="_blank" rel="noreferrer"> 
      <img src={svelteLogo} class="logo svelte" alt="Svelte Logo" />
    </a>
  </div>
  <h1>SvEcho</h1>

  <div class="card">
    <Counter />
  </div>

  <p>
    Check out <a href="https://github.com/sveltejs/kit#readme" target="_blank" rel="noreferrer">SvelteKit</a>, the official Svelte app framework powered by Vite!
  </p>
  {#if data}
    <h3>{data.message}</h3>
  {/if}
  
</main>

<style>
  .logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
    transition: filter 300ms;
  }
  .logo:hover {
    filter: drop-shadow(0 0 2em #646cffaa);
  }
  .logo.svelte:hover {
    filter: drop-shadow(0 0 2em #ff3e00aa);
  }
  .read-the-docs {
    color: #888;
  }
</style>