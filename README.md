# lv.github.io
_Repository hosting my personal blog: [luis.vi](https://luis.vi/)_

## Content
Every Markdown file in `web/content/` becomes a page. The path stays the same and `.md` becomes `.html`, so `blog/first.md` is built as `blog/first.html`. A new file is picked up automatically; nothing else needs editing.

`web/content/index.md` is required: it's the home page, and the build fails without it.

The site-wide title, subtitle and author are set in `configs/site.json`.
