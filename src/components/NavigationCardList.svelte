<ol>
    <slot />
</ol>

<style>
    ol {
        display: grid;
        padding: 0;
        margin-bottom: 0 !important;
        gap: var(--spacing-md);
        grid-template-columns: repeat(auto-fit, minmax(15rem, 1fr));
        list-style: none;
    }
</style>
