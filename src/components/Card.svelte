<script lang="ts">
    export let href: string;

    export let iconPath: string | undefined = undefined;
    export let centerSlottedContent: boolean = false;
    export let padding: string = "var(--spacing-md)"
    export let boxShadowSize: string = "var(--spacing-2xs)";

    let isExternalLink = href.startsWith("http");
</script>

<li>
    <a
        {href}
        class={`${iconPath ? "icon" : ""} ${centerSlottedContent ? "center-slotted-content" : ""}`}
        style="
            --icon-path: url({iconPath});
            --card-box-shadow-size: {boxShadowSize};
            --card-padding: {padding};
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
        padding: var(--card-padding);
        height: 100%;
        display: block;
        text-decoration: none;
        background-color: var(--color-surface);
        transition: all 0.15s;
        border: var(--border-md);

        &.center-slotted-content {
            display: flex;
            justify-content: center;
            align-items: center;
            aspect-ratio: 1 / 1;
        }

        &:hover p {
            color: var(--color-background);
        }

        & p {
            font-size: var(--body-sm);
            font-weight: 300;
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
            margin-right: calc(0.5 * var(--rem));
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
            box-shadow: var(--card-box-shadow-size) var(--card-box-shadow-size) var(--color-foreground);
        }

        a:hover {
            transform: translate(var(--card-box-shadow-size), var(--card-box-shadow-size));
        }
    }

    @media (prefers-color-scheme: dark) and (prefers-reduced-motion: no-preference) {
        a:hover {
            transform: translateY(-3px);
        }
    }
</style>
