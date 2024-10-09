<script lang="ts">
    import type { CollectionEntry } from "astro:content";
    import NoteNavigationCard from "@components/NoteNavigationCard.svelte";

    export let notes: CollectionEntry<"notes">[];
    const groupedNotes: {[key: string]: CollectionEntry<"notes">[]} = {};

    notes.forEach(note => {
        let firstLetter = note.slug.charAt(0).toLowerCase();
        if (!groupedNotes[firstLetter]) {
            groupedNotes[firstLetter] = [];
        }
        groupedNotes[firstLetter].push(note);
    });
</script>

{#each Object.entries(groupedNotes) as [firstLetter, notes]}
    <h2 id={firstLetter}>{firstLetter.toUpperCase()}</h2>
    <ol>
        {#each notes as note}
            <li><NoteNavigationCard {...note.data} {...note} /></li>
        {/each}
    </ol>
{/each}

<style>
    ol {
        display: grid;
        padding: 0;
        gap: var(--spacing-md);
        grid-template-columns: repeat(auto-fit, minmax(15rem, 1fr));
        list-style: none;
    }
</style>
