<script lang="ts">
    import BreadcrumbNavigationItem from "@components/BreadcrumbNavigationItem.svelte";

    export let currentPagePath: string;
    const breadcrumbItems = [{ href: "/", displayText: "th7mo" }];

    currentPagePath.split("/").forEach((pathPart, index, array) => {
        const href = `/${array.slice(0, index + 1).join("/")}`;
        const displayText = pathPart.replaceAll("-", " ");
        breadcrumbItems.push({ href, displayText });
    });
</script>

<header>
    <nav>
        <ol>
            {#each breadcrumbItems as item}
                <BreadcrumbNavigationItem {...item} />
            {/each}
        </ol>
    </nav>
</header>

<style>
    header {
        top: 0;
        position: sticky;
        padding: var(--spacing-md) 0;
        margin: var(--layout-xl) 0;
        background-color: var(--color-background);
        z-index: 1;
    }

    ol {
        display: flex;
        padding: 0;
        flex-wrap: wrap;
        list-style: none;
        user-select: none;
        -webkit-user-select: none;
        font-size: var(--body-lg);
    }
</style>
