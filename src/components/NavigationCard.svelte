<script lang="ts">
    export let title: string;
    export let description: string | undefined = undefined;
    export let href: string;
    export let iconPath: string = "/images/note.svg";
    if (href.startsWith("http")) {
        iconPath = "/images/external-link.svg";
    }
</script>

<li>
    <a class="navigation-card" {href} style="--icon-path: url({iconPath})">
        {title}
        {#if description}
            <p>{description}</p>
        {/if}
    </a>
</li>

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

        &::before {
            content: "";
            background-color: var(--color-foreground);
            width: var(--body-md);
            height: var(--body-md);
            display: inline-block;
            margin-right: 0.5rem;
            mask: var(--icon-path);
            -webkit-mask: var(--icon-path);
            transform: translateY(2px);
        }

        &::after {
            display: none !important;
        }

        &:hover::before {
            background-color: var(--color-background);
        }
    }
</style>
