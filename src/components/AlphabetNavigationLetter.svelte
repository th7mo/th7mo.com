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

    @media (prefers-color-scheme: light) and (prefers-reduced-motion: no-preference) {
        a {
            --shadow-size: 0.15rem;
            box-shadow: var(--shadow-size) var(--shadow-size) var(--color-foreground);
        }

        a:hover {
            transform: translate(var(--shadow-size), var(--shadow-size));
        }
    }

    @media (prefers-color-scheme: dark) and (prefers-reduced-motion: no-preference) {
        a:hover {
            transform: translateY(-2px);
        }
    }
</style>
