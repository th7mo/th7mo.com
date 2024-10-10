<script lang="ts">
    export let letter: string;
    export let href: string;
</script>

<li>
    <a class="navigation-card" {href}>
        {letter.toUpperCase()}
    </a>
</li>

<style>
    a {
        display: flex;
        justify-content: center;
        align-items: center;
        width: var(--alphabet-navigation-letter-size);
        height: var(--alphabet-navigation-letter-size);
    }

    @media (prefers-color-scheme: light) {
        a {
            box-shadow: 0.15rem 0.15rem var(--color-foreground);
        }
    }

    @media (prefers-color-scheme: light) and (prefers-reduced-motion: no-preference) {
        a:hover {
            transform: translate(0.15rem, 0.15rem);
        }
    }

    @media (prefers-color-scheme: dark) and (prefers-reduced-motion: no-preference) {
        a:hover {
            transform: translateY(-2px);
        }
    }
</style>
