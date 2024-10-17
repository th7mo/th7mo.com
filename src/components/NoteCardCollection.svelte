<script lang="ts">
    import type { CollectionEntry } from "astro:content";
    import TerminalList from "@components/TerminalList.svelte";
    import TerminalListItem from "@components/TerminalListItem.svelte";

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
    <TerminalList>
        {#each notes as note}
            <TerminalListItem short date={note.data.dateLastModified}>
                <a href={`${currentPagePath}/${note.slug}`}>{note.data.title}</a>
            </TerminalListItem>
        {/each}
    </TerminalList>
{/each}
