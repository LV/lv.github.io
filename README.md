# lv.github.io
_Repository hosting my personal blog: [luis.vi](https://luis.vi/)_

The site is built by a small custom static site generator written in Go, using only the standard library. The repository follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

## Prerequisites
- Go (the version in `go.mod`)
- `make`

## Make targets
- `make` / `make serve` — build the site and preview it at http://localhost:8080
- `make build` — generate the site into `_site/`
- `make clean` — delete `_site/`
- `make fmt` — format the Go code

## Layout
- `cmd/sitegen` — the generator
- `cmd/serve` — the local preview server
- `internal/site` — paths shared by both programs
- `web/template/layout` — the layout every page is rendered into
- `web/template/page` — one template per page; each path mirrors its output path in `_site/`
- `web/static` — CSS, images, and other files written into `_site/` unchanged
- `drafts` — unpublished writing

## Deployment
Pushing to `master` runs `.github/workflows/static.yml`, which builds the site with `make build` and deploys `_site/` to GitHub Pages. It only runs when a commit touches the site's sources. Pull requests into `master` run `.github/workflows/check.yml`, which checks formatting, `go.mod` tidiness, `go vet`, tests, and that the site builds.
