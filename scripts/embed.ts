#!/usr/bin/env -S deno run -A

import { resolve } from "@std/path";

const rootDir = resolve(import.meta.dirname!, "../");
const dataDir = resolve(rootDir, "data");
const embedFilePath = resolve(rootDir, "embed.json");

const embed: Record<string, unknown> = {};

for (const file of Deno.readDirSync(dataDir)) {
  embed[file.name] = Array.from(Deno.readFileSync(resolve(dataDir, file.name)));
}

Deno.writeTextFileSync(embedFilePath, JSON.stringify(embed));
