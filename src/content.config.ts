import { defineCollection, z } from "astro:content";
import { glob } from "astro/loaders";

const notes = defineCollection({
  loader: glob({ pattern: "**/*.md", base: "./src/content/notes" }),
  schema: z.object({
    title: z.string(),
    description: z.string(),
    isPublic: z.boolean(),
    dateCreated: z.string(),
    dateLastModified: z.string(),
  }),
});

export const collections = { notes };
