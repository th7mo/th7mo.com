<script lang="ts">
    export let dateCreated: string;
</script>

<footer>Posted on <time>{dateCreated}</time></footer>

<style>
    footer {
        margin: 0 auto;
        max-width: var(--page-width);
        padding: var(--page-padding);
        text-align: end;
        width: 100%;
    }
</style>
