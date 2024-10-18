<script lang="ts">
    export let href: string;
    export let displayText: string;
</script>

<li><a {href}>{displayText}</a></li>

<style>
    li {
        & a {
            font-weight: var(--weight-normal);
            text-decoration: none;

            &:hover {
                text-decoration-thickness: 1.2px;
            }
        }

        &:last-child a {
            font-weight: var(--weight-bold);
        }

        &:not(:last-child)::after {
            padding: 0 var(--spacing-sm);
            content: "/";
        }
    }
</style>
