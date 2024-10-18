<script lang="ts">
    export let dir = false;
    export let short = false;
    export let executable = false;
    export let group = "th7mo";
    export let date: string;

    while (group.length < 5) {
        group += " ";
    }

    let fileType = dir ? "d": "-";
    let ownerPermissions = "rw" + (executable ? "x" : "-");
    let guestPermissions = "r-" + (executable ? "x" : "-");
    let permissions = fileType + ownerPermissions + guestPermissions + guestPermissions;
</script>

{#if short}
    <li><time>{date}</time> &nbsp;<slot /></li>
{:else}
    <li>{permissions} &nbsp;<slot /></li>
{/if}
