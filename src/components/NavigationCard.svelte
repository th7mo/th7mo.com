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

        &:hover::before {
            background-color: var(--color-background);
        }

        &::after {
            display: none !important;
        }
    }

    @media (prefers-color-scheme: light) and (prefers-reduced-motion: no-preference) {
        a {
            --shadow-size: var(--spacing-2xs);
            box-shadow: var(--shadow-size) var(--shadow-size) var(--color-foreground);
        }

        a:hover {
            transform: translate(var(--shadow-size), var(--shadow-size));
        }
    }

    @media (prefers-color-scheme: dark) and (prefers-reduced-motion: no-preference) {
        a:hover {
            transform: translateY(-3px);
        }
    }
</style>
