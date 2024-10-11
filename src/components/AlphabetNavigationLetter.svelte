<script lang="ts">
    import Card from "@components/Card.svelte";
    export let letter: string;
    export let href: string;
</script>

<Card {href} boxShadowSize="0.15rem" padding="var(--spacing-xs)" centerSlottedContent>
    {letter.toUpperCase()}
</Card>
