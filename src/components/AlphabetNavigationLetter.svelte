<script lang="ts">
    import Card from "@components/Card.svelte";
    export let letter: string;
    export let href: string;
</script>

<Card {href} shadowSize="0.15rem" spacing="var(--spacing-xs)" center>
    {letter.toUpperCase()}
</Card>
