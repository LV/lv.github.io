# lv.github.io
_Repository hosting my personal blog: [luis.vi](https://luis.vi/)_

It follows the [Plain Vanilla Web guildelines](https://plainvanillaweb.com/pages/sites.html).

## Pre-requisites
- `npm` for managing dependencies and building the project
- `python3` for running scripts

## Available scripts
- `npm run dev` — serve the `public` directory locally with caching disabled, for development
- `npm run serve` — serve the `public` directory locally the way GitHub Pages does (see below)
- `npm run check` — type-check the project with `tsc`
- `npm run feedgen` — generate the RSS/Atom feed
- `npm run build` — run `check` then `feedgen`

### Matching GitHub Pages
`npm run serve` runs `http-server` with flags that mimic how GitHub Pages serves the site:
- `-c600` caches responses for 10 minutes. GitHub doesn't document this value, but it's the `cache-control: max-age=600` header GitHub Pages sends (checked 2026-10-15). To re-check it, run `curl -sI https://luis.vi/ | grep -i cache-control`
- `-d false` turns off directory listings. For a directory with no `index.html`, GitHub Pages returns a 404, while `http-server` returns a 403
