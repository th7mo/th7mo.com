<script lang="ts">
    import type { CollectionEntry } from "astro:content";
    import AlphabetNavigationLetter from "@components/AlphabetNavigationLetter.svelte";

    export let notes: CollectionEntry<"notes">[];  

    const letters = [...new Set(notes.map(note => note.slug.charAt(0).toLowerCase()))];
</script>

<ol>
    {#each letters as letter}
        <AlphabetNavigationLetter {letter} />
    {/each}
</ol>

<style>
    ol {
        margin: var(--spacing-lg) 0 calc(2 * var(--spacing-lg)) 0;
        padding: 0;
        display: grid;
        gap: var(--spacing-sm);
        grid-template-columns: repeat(auto-fill, var(--alphabet-navigation-letter-size));
        justify-content: space-between;
        list-style: none;
    }
</style>
