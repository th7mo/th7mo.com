<script lang="ts">
    export let letter: string;
    export let href: string;
</script>

<li>
    <a class="navigation-card" href={href}>
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
</style>
