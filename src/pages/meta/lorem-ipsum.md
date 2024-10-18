---
layout: "@layouts/MarkdownLayout.astro"
title: "Lorem Ipsum"
description: "The sandbox page for th7mo.com"
---

It is used for this website to stress-test its features in a sandbox
environment. Lorem Ipsum is placeholder text for previewing layouts.

## Body text

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin suscipit nisl
ut ante porta, sed egestas eros rhoncus. Nunc semper tortor odio, sed porta
ipsum tincidunt imperdiet. Praesent turpis ipsum, finibus eu porta sit amet,
condimentum non elit. Maecenas fringilla scelerisque lacus, vel tempus odio
aliquam sit amet. Nullam scelerisque facilisis dui, at maximus libero sodales
vel. Cras cursus libero dolor, ac pharetra lectus blandit fermentum. Quisque
arcu diam, mollis sed consequat at, lacinia at urna.

Nulla accumsan odio et arcu semper, in faucibus enim varius. Nullam mollis
sapien dui, id eleifend dolor feugiat at. Integer pellentesque leo sit amet
pretium dignissim. Nulla facilisis ultricies quam, ut ullamcorper elit laoreet
vitae. Sed diam sem, lobortis in massa eu, mattis congue orci. Sed viverra
ipsum ut mauris posuere, non tincidunt justo lobortis.

Donec ut quam sed ipsum
condimentum efficitur. Nullam sit amet nunc sed ligula mollis pretium. Sed
porttitor ullamcorper sagittis. Nunc ullamcorper gravida nunc at laoreet. Fusce
sed nibh tortor. Suspendisse erat libero, finibus ut eros nec, auctor porttitor
arcu. Etiam arcu sapien, iaculis non dolor at, blandit vestibulum nunc. Morbi
rhoncus eros et risus fringilla, vel dignissim urna mattis. Vestibulum fringilla
metus nec turpis tempus, sit amet interdum nulla imperdiet. Curabitur et congue
tortor.

Ut a pharetra nibh, non aliquam nulla. Aliquam dapibus tempor ipsum, ultrices
sagittis lorem volutpat non. Nulla facilisi. Donec vel iaculis risus, non
consectetur arcu. Maecenas enim diam, placerat nec dui sed, convallis congue
sem. Morbi vel elementum ex. Nullam porttitor libero at nunc sagittis malesuada.
Donec pharetra augue vestibulum leo viverra ullamcorper non quis nibh.
Suspendisse malesuada vehicula arcu, ut mollis turpis rutrum a.

## Inline formatting

*Text with italics*

**Text that is strong**

***Italic and strong***

***~~Italic, strong and strikethrough~~***

~~strikethrough~~

Some<sup>sup</sup> text

Some<sub>sub</sub> text.

`Inline code` `with` `multiple` `after` `each other`

Use the `<p>` tag for regular paragraphs.

<kbd>Keyboard</kbd> <kbd>Input</kbd> <kbd>blocks</kbd>

Press <kbd>Ctrl</kbd> + <kbd>Shift</kbd> + <kbd>R</kbd> to refresh a page.


A footnote[^1]

A double footnote[^1] [^2].

[^1]: This is an example footnote.
[^2]: This is also an example footnote.

A horizontal rule:

---

Two horizontal rules:

---

---

## Links

[External Link](https://google.com)

[Internal link](color-palette)

[Internal `code` link](color-palette)

[Internal **bold** link](color-palette)

[Internal ~~strikethrough~~ link](color-palette)

[Internal<sup>sup</sup> link](color-palette)

[Internal <kbd>Input</kbd> link](color-palette)

[Section link](#lists)

## Blockquotes

A regular blockquote:

> "Meditation is a useful tool – but we actually need to stop the people who are
> pouring itching powder on us."

A blockquote with a code block inside:

> ```sh
> sudo apt install discord
> ```

Nested blockquotes:


> "Meditation is a useful tool – but we actually need to stop the people who are
> pouring itching powder on us."
>
> > "Meditation is a useful tool – but we actually need to stop the people who are
> > pouring itching powder on us."

A triple nested blockquote:

> > > "Meditation is a useful tool – but we actually need to stop the people who are
> > > pouring itching powder on us."

A blockquote with a list inside:

> * level 1
>     * level 2
>         * level 3
>             * level 4
>                 * level 5

## Lists

A normal ordered list:

1. Type <kbd>i</kbd> to go into insert mode.
2. Write what you want to write.
3. Press <kbd>esc</kbd> to return to normal mode.
4. Enter the `:wq` command to write the file and exit.

A normal unordered list:

* Apples
* Bananas
* Oranges
* Kiwi's

Regular list with other elements:

1. with `inline code`
2. and some code blocks:
   ```css
   p {
       color: var(--color-foreground);
       background-color: var(--color-background);
   }
   ```
3. and a blockquote:
   > "Here is an inspiring quote"

Heavy indented unordered list:

* level 1
    * level 2
    * level 2.2
        * level 3
        * level 3.2
            * level 4
    * level 2.3
        * level 3.3

Heavy indented ordered list:
  
1. level 1
    1. level 2
        1. level 3
        2. Level 3.2
            1. level 4

## Code blocks

Shell:

```sh
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
```

```sh
sudo apt update && sudo apt upgrade
```

```sh
sudo apt update a long line that wrappes around to the next line if it is even longer than this
```

Web development:

```javascript
if (number === 1) {
    console.log("1 is neither prime nor composite number.");
}

// check if number is greater than 1
else if (number > 1) {

    // looping through 2 to number/2
    for (let i = 2; i <= number / 2; i++) {
        if (number % i == 0) {
            isPrime = false;
            break;
        }
    }

    if (isPrime) {
        console.log(`${number} is a prime number`);
    } else {
        console.log(`${number} is a not prime number`);
    }
}
```

```css
@media (prefers-color-scheme:dark) {
    :root {
        --color-foreground: #eee;
        --color-background: #000;
        --color-surface: #1f1e25;
    }
}
```

```svelte
<script lang="ts">
    export let title: string;
    export let description: string;
</script>

<head>
    <meta charset="UTF-8">
    <meta name="author" content="Thimo van Velzen">
    <meta name="keywords" content="website">
    <meta name="description" content={description}>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>{title}</title>
</head>
```

## Tables

A small table:

| A | B | C |
|---|---|---|
| 1 | 2 | 3 |
| 4 | 5 | 6 |
| 7 | 8 | 9 |
 
A regular table:

| Column 1                    | Column 2            | Column 3  |
|-----------------------------|---------------------|-----------|
| Here is some data for col 1 | Some more for col 2 | and col 3 |
| Here is some data for col 1 | Some more for col 2 | and col 3 |
| Here is some data for col 1 | Some more for col 2 | and col 3 |
| Here is some data for col 1 | Some more for col 2 | and col 3 |
 
A table with inline code-blocks:

| Command                             | Action                                           |
|-------------------------------------|--------------------------------------------------|
| `npm install`                       | Installs dependencies                            |
| `npm run dev`                       | Starts local dev server at `localhost:4321`      |
| `npm run preview`                   | Preview your build locally, before deploying     |
| `npm run build`                     | Build your production site to `./dist/`          |
| `npm run astro ...`                 | Run CLI commands like `astro add`, `astro check` |
| `npm run astro -- --help`           | Get help using the Astro CLI                     |

A big table:

| This is a really long table with possible wraps to new line because the table heading is too big | This is a really long table with possible wraps to new line because the table heading is too big |
|--------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------|
| Lorem ipsum                                                                                      | idem dito                                                                                        |
| More values                                                                                      | This is a really long table with possible wraps to new line because the table heading is too big |
 
A table with formatting:

| column   | column |
|----------|--------|
| **test** | *test* |

A table with `sub` and `sup` text:

| column             | column             |
|--------------------|--------------------|
| some<sub>sub</sub> | some<sup>sup</sup> |


