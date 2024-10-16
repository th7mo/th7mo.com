<script lang="ts">
    export let title: string;
    export let description: string;
</script>

<head>
    <meta charset="UTF-8">
    <meta name="author" content="Thimo van Velzen">
    <meta name="keywords" content="HTML, CSS, Astro, Svelte, wiki, minimalist, simple, notes, note-taking">
    <meta name="description" content={description}>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>{title}</title>
    <link rel="icon" href="/images/note.svg">
    <link
        rel="preload"
        href="/fonts/averia-300-normal.woff2"
        as="font"
        type="font/woff2"
        crossorigin="anonymous"
    >
    <link
        rel="preload"
        href="/fonts/averia-700-normal.woff2"
        as="font"
        type="font/woff2"
        crossorigin="anonymous"
    >
    <link
        rel="preload"
        href="/fonts/inconsolata-mono.woff2"
        as="font"
        type="font/woff2"
        crossorigin="anonymous"
    >
</head>

