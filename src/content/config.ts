import { z, defineCollection } from "astro:content";

const notes = defineCollection({
    type: 'content',
    schema: z.object({
      title: z.string(),
      description: z.string(),
      isPublic: z.boolean()
    })
});

export const collections = { notes };
