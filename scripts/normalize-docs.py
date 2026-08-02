#!/usr/bin/env python3
"""Normalise catonetworks/cato Terraform registry docs for upjet's scraper.

Upjet's scraper locates a document's prelude with the XPath expression
``//text()[contains(., "description") and contains(., "page_title")]`` and its
import statements with ``//code[@class="language-shell"]/text()``. A large part
of the catonetworks/cato docs carry neither a ``page_title`` nor a
``description`` key in their front matter, and their import examples sit in
unlabelled code fences, so the scraper aborts on the first such file.

This script rewrites the docs in place: it fills in the missing front matter
keys (deriving the description from the first paragraph of the body) and labels
the opening fence of every code block in the "## Import" section as ``shell``.
It is idempotent, so re-running it over already normalised docs is a no-op.

Usage: normalize-docs.py <docs/resources directory>
"""

import os
import re
import sys

PREFIX = "cato_"
SKIP_PARAGRAPH_PREFIXES = ("#", "```", "!>", "->", "~>", "|", "---")


def first_paragraph(body: str) -> str:
    for block in re.split(r"\n\s*\n", body):
        block = block.strip()
        if not block or block.startswith(SKIP_PARAGRAPH_PREFIXES):
            continue
        return re.sub(r"\s+", " ", block)[:400]
    return ""


def label_import_fences(body: str) -> str:
    def rewrite(match: "re.Match[str]") -> str:
        head, rest = match.group(1), match.group(2)
        lines, inside = [], False
        for line in rest.split("\n"):
            if re.fullmatch(r"```\s*", line):
                lines.append("```" if inside else "```shell")
                inside = not inside
            else:
                if line.startswith("```"):
                    inside = not inside
                lines.append(line)
        return head + "\n".join(lines)

    return re.sub(r"(?s)(##\s+Import\s*\n)(.*)$", rewrite, body)


def normalize(path: str, filename: str) -> bool:
    with open(path, encoding="utf-8") as handle:
        text = handle.read()

    match = re.match(r"^---\n(.*?)\n---\n(.*)$", text, re.S)
    front_matter, body = (match.group(1), match.group(2)) if match else ("", text)

    lines = [line for line in front_matter.split("\n") if line.strip()]
    keys = {line.split(":", 1)[0].strip() for line in lines if ":" in line}
    resource = PREFIX + filename[: -len(".md")]

    if "page_title" not in keys:
        lines.append(f'page_title: "{resource} Resource - terraform-provider-cato"')
    if "subcategory" not in keys:
        lines.append('subcategory: ""')
    if "description" not in keys:
        lines.append("description: |-\n  " + first_paragraph(body).replace('"', "'"))

    normalized = "---\n" + "\n".join(lines) + "\n---\n" + label_import_fences(body)
    if normalized == text:
        return False
    with open(path, "w", encoding="utf-8") as handle:
        handle.write(normalized)
    return True


def main() -> int:
    if len(sys.argv) != 2:
        print(f"usage: {os.path.basename(sys.argv[0])} <docs directory>", file=sys.stderr)
        return 2
    docs_dir = sys.argv[1]
    if not os.path.isdir(docs_dir):
        print(f"not a directory: {docs_dir}", file=sys.stderr)
        return 1

    changed = 0
    for filename in sorted(os.listdir(docs_dir)):
        if filename.endswith(".md") and normalize(os.path.join(docs_dir, filename), filename):
            changed += 1
    print(f"normalised {changed} document(s) in {docs_dir}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
