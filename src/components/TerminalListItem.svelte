<script lang="ts">
    export let dir = false;
    export let executable = false;
    export let date: string | undefined = undefined;

    let fileType = dir ? "d": "-";
    let ownerPermissions = "rw" + (executable ? "x" : "-");
    let guestPermissions = "r-" + (executable ? "x" : "-");
    let permissions = fileType + ownerPermissions + guestPermissions + guestPermissions;
</script>

{#if date}
    <li><time>{date}</time> &nbsp;<slot /></li>
{:else}
    <li>{permissions} &nbsp;<slot /></li>
{/if}
