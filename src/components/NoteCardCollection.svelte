<script lang="ts">
    import type { CollectionEntry } from "astro:content";
    import CardCollection from "@components/CardCollection.svelte";
    import NoteNavigationCard from "@components/NoteNavigationCard.svelte";

    export let currentPagePath: string;
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
    <CardCollection>
        {#each notes as note}
            <NoteNavigationCard href={`${currentPagePath}/${note.slug}`} {...note.data} />
        {/each}
    </CardCollection>
{/each}
