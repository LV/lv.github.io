# lv.github.io
_Repository hosting my personal blog: [luis.vi](https://luis.vi/)_

## Content
Every Markdown file in `web/content/` becomes a page. The path stays the same and `.md` becomes `.html`, so `blog/first.md` is built as `blog/first.html`. A new file is picked up automatically; nothing else needs editing.

The site-wide title, subtitle and author are set in `configs/site.json`.

## Required files
The build checks that these files exist before doing anything else. If any are missing, it stops and names every one of them.
- `configs/site.json`: the site-wide settings
- `web/content/index.md`: the home page
- `web/template/base.html`: the layout every page is rendered into
- `web/static/favicon.ico`: the favicon (light mode)
- `web/static/favicon-dark.ico`: the favicon (dark mode)
- `web/static/index.css`: the stylesheet
