<script lang="ts">
    import Card from "@components/Card.svelte";

    export let title: string;
    export let description: string | undefined = undefined;
    export let href: string;
    export let iconPath: string = "/images/note.svg";

    let isExternalLink = href.startsWith("http");
    if (isExternalLink) {
        iconPath = "/images/external-link.svg";
    }
</script>

<Card {href} {iconPath} boxShadowSize="var(--spacing-2xs)">
    {title}
    {#if description}
        <p>{description}</p>
    {/if}
</Card>
