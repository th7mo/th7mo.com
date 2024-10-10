<script lang="ts">
    export let title: string;
    export let description: string;
    export let href: string;
</script>

<a class="navigation-card" href={href}>
    {title}
    <p>{description}</p>
</a>

<style>
    a {
        font-size: var(--body-lg);
        padding: var(--spacing-md);
        display: block;
        height: 100%;

        &:hover p {
            color: var(--color-background);
        }

        & p {
            font-size: var(--body-sm);
            font-weight: 400;
            hyphens: none;
        }
    }
</style>
