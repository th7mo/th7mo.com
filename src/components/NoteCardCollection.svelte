<script lang="ts">
    import type { CollectionEntry } from "astro:content";
    import NavigationCardList from "@components/NavigationCardList.svelte";
    import NavigationCard from "@components/NavigationCard.svelte";

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
    <NavigationCardList>
        {#each notes as note}
            <NavigationCard href={`${currentPagePath}/${note.slug}`} {...note.data} />
        {/each}
    </NavigationCardList>
{/each}
