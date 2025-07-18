"""Luis Victoria's Atom Feed Generator"""
FEED_GENERATOR_VERSION: str = "0.1.0"

# IDEA: Use Git `master` metadata to add timestamps and stuff to blogs and entries
# HELPFUL LINKS:
# Guide and Tips: https://kevincox.ca/2022/05/06/rss-feed-best-practices/ as a guide
# W3 Concise Guide: https://validator.w3.org/feed/docs/atom.html
# Official Spec: https://datatracker.ietf.org/doc/html/rfc4287

from dataclasses import dataclass
from datetime import datetime
from pathlib import Path

# CONSTANTS
AUTHOR_NAME: str = "Luis Victoria"
AUTHOR_EMAIL: str = "v@luis.vi"
DOMAIN_NAME: str = "luis.vi"
BASE_URL: str = "https://" + DOMAIN_NAME
FEED_TITLE: str = AUTHOR_NAME
FEED_SUBTITLE: str = "Creating to Understand"
SITE_DIR: Path = Path(__file__).resolve().parent.parent / "public" # TODO: Clean out things in `.gitignore`?
FEED_FILE: Path = SITE_DIR / "feed.atom"


@dataclass
class Entry:
    """An RSS entry."""
    title: str
    url: str
    permalink_id: str
    published: datetime
    last_updated: datetime
    html_content: str


@dataclass
class FeedContent:
    feed_last_updated: datetime
    entries: list[Entry]


def get_feed_content() -> FeedContent:
    """Parse `feed.atom`, `public/`, and Git metadata to generate `FeedContent`."""
    # STUB: Parse content
    return FeedContent(
        feed_last_updated=datetime.now(),
        entries=[],
    )


def generate_file_contents(feed: FeedContent) -> str:
    """Generates the entire `feed.atom` file as a string."""

    # TODO: Fix the `<id>` tag under the `<updated>` tag
    final_str = ""
    final_str += f"""<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns="http://www.w3.org/2005/Atom">
  <title type="text">{FEED_TITLE}</title>
  <subtitle type="text">{FEED_SUBTITLE}</subtitle>
  <updated>{feed.feed_last_updated.isoformat()}</updated>
  <id>tag:{DOMAIN_NAME},2024</id>
  <link rel="alternate" type="text/html" href="{BASE_URL}"/>
  <link rel="self" type="application/atom+xml" href="{BASE_URL+"/feed.atom"}"/>
  <rights>© {str(datetime.now().year)} {AUTHOR_NAME}. All rights reserved.</rights>
  <generator uri="https://github.com/LV/lv.github.io/blob/master/scripts/feedgen.py" version="{FEED_GENERATOR_VERSION}">
    Luis Victoria's Atom Feed Generator
  </generator>
  <author>
    <name>{AUTHOR_NAME}</name>
    <email>{AUTHOR_EMAIL}</email>
    <uri>{BASE_URL}</uri>
  </author>
  <logo>/favicon.ico</logo>
"""

    # TODO: Generate entries
    # Check for entries in `SITE_DIR`, and generate them accordingly

    final_str += "</feed>\n"
    return final_str


def write_file(content: str) -> None:
    """Writes the `feed.atom` file."""
    FEED_FILE.write_text(content) # overwrites contents if it exists


def main() -> None:
    """Script entrypoint."""
    write_file(generate_file_contents(get_feed_content()))


if __name__ == "__main__":
    main()
