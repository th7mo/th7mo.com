<script lang="ts">
    export let href: string;

    export let iconPath: string | undefined = undefined;
    export let center: boolean = false;
    export let spacing: string = "var(--spacing-md)"
    export let shadowSize: string = "";

    let isExternalLink = href.startsWith("http");
</script>

<li>
    <a
        {href}
        class={`${iconPath ? "icon" : ""} ${center ? "center" : ""}`}
        style="
            --icon-path: url({iconPath});
            --card-shadow-size: {shadowSize};
            --card-spacing: {spacing};
        "
        target={isExternalLink ? "_blank" : undefined}
        rel={isExternalLink ? "noopener noreferrer" : undefined}
    >
        <slot />
    </a>
</li>

<style>
    a {
        font-size: var(--body-lg);
        padding: var(--card-spacing, var(--spacing-md));
        height: 100%;
        display: block;
        text-decoration: none;
        background-color: var(--color-surface);
        transition: all 0.15s;
        border: var(--border-md);

        &.center {
            display: flex;
            justify-content: center;
            align-items: center;
        }

        &:hover p {
            color: var(--color-background);
        }

        & p {
            font-size: var(--body-sm);
            font-weight: 400;
            hyphens: none;
        }

        &::before {
            display: none;
        }

        &.icon::before {
            content: "";
            display: inline-block;
            background-color: var(--color-foreground);
            width: var(--body-md);
            height: var(--body-md);
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
            box-shadow: var(--card-shadow-size) var(--card-shadow-size) var(--color-foreground);
        }

        a:hover {
            transform: translate(var(--card-shadow-size), var(--card-shadow-size));
        }
    }

    @media (prefers-color-scheme: dark) and (prefers-reduced-motion: no-preference) {
        a:hover {
            transform: translateY(-3px);
        }
    }
</style>
