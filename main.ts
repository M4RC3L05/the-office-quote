import { parse } from "@std/yaml/parse";
import { parseArgs } from "@std/cli";
import meta from "./deno.json" with { type: "json" };

const help = `
The Office Quote

Display a random quote from the Office US

Usage: the-office-quote [OPTIONS]

Options:
  --help, -h      Display this help menu
  --version, -V   Display version
`.trim();

const quotesFile = new URL("./data/quotes.yaml", import.meta.url);
const quotesFileContents = Deno.readTextFileSync(quotesFile);

const getRandomInt = (min: number, max: number) => {
  min = Math.ceil(min);
  max = Math.floor(max);

  return Math.floor(Math.random() * (max - min + 1)) + min;
};

const printHelp = () => {
  console.log(help);
};

const printVersion = () => {
  console.log(`v${meta.version}`);
};

const printQuote = () => {
  const quotes = parse(quotesFileContents) as { text: string }[];

  if (!Array.isArray(quotes)) {
    console.error(`Could not parse quotes`);
    Deno.exit(1);
  }

  const quoteIndex = getRandomInt(0, quotes.length - 1);
  const quote = quotes.at(quoteIndex);

  if (!quote) {
    console.error(`Could not get quote for index "${quoteIndex}"`);
    Deno.exit(1);
  }

  console.log(quote.text);
};

if (import.meta.main) {
  const args = parseArgs(Deno.args, {
    alias: { help: "h", version: "V" },
    boolean: ["help", "version"],
  });

  if (args.h || args.help) {
    printHelp();
    Deno.exit(0);
  }

  if (args.version || args.V) {
    printVersion();
    Deno.exit(0);
  }

  await printQuote();
}
