---
layout: "@layouts/MarkdownLayout.astro"
title: "Design choices"
description: "Design choices that made the th7mo.com website what it is today"
---

If this website is more readable than
[Reading Mode](https://support.google.com/chrome/answer/14218344) implemented
on most modern browsers, then I've achieved my goal.

## No Table Of Contents (TOC)

I chose to *not* include a Table Of Contents at the start of each note. This is
because I want to keep my notes [atomic](/notes/atomic). If the note is long enough
to deserve a Table Of Contents, it is probably covering multiple topics and
could be split up into separate notes.

## No website versions

Although I think [semantic versioning](/notes/semantic-versioning) is great to keep
track of changes in a project, I chose to *not* keep track of versions for
[th7mo.com](https://th7mo.com). I am the only one maintaining and developing
the website. Versioning changes made to the website is not useful to anyone.

## No JavaScript

This website is only for reading content and nothing more. I try to keep the
website simplistic, lightweight and optimized for its purpose: serving
read-only content. I use [Astro](/notes/astro) to convert
[Markdown](/notes/markdown) files to static, plain HTML and [CSS](/notes/css).

## Not tagging notes

Tagging notes can be a great and flexible system for adding metadata to notes.
However, I chose not to implement it for my collection because I believe the
pros don't outweigh the cons. This decision keeps complexity out of the website.

## Not nesting notes

All the notes are on the same root level. Categorizing the notes would be a
waste of time because the notes have a complex relationship with each other
that could be restricted by categorizing them into directories.

## Content as a submodule

The content for the website is tracked in a different [Git](/notes/git)
repository as the website. I have two views for the content, one in
[Obsidian](/notes/obsidian) and one is the website. In the future I might add
even more or change existing ones. I can edit my notes without knowing the
implementation details for presenting the content.

## Fonts

The website does not use custom web fonts, except for monospaced text. Web fonts
are designed to be readable to the wider audience and works perfectly fine. I
chose to have a custom monospaced font to make code blocks more readable on
mobile. [sustainable.dev](https://the-sustainable.dev/do-you-really-need-that-custom-webfont/)
explains this in more detail.

## No images

I want to keep my content repository manageable over the long term. Including
images from the start could lead to a significant increase in the Git
repository's size, which may cause storage and usability issues down the line.

## Whitespace below the main content

Each page has equal whitespace above and below the main content. This additional
whitespace at the bottom of each page ensures that the last line of text is
positioned closer to the center of the screen rather than at the bottom,
enhancing the overall reading experience.
