"""Luis Victoria's Atom Feed Generator"""
FEED_GENERATOR_VERSION: str = "0.1.0"

# IDEA: Use Git `master` metadata to add timestamps and stuff to blogs and entries
# Using https://kevincox.ca/2022/05/06/rss-feed-best-practices/ as a guide

from dataclasses import dataclass
from datetime import datetime
from pathlib import Path

# CONSTANTS
AUTHOR_NAME: str = "Luis Victoria"
BASE_URL: str = "https://luis.vi"
FEED_NAME: str = "Luis Victoria's Blog"
SITE_DIR: Path = Path(__file__).resolve().parent.parent / "public"
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
    base_url: str
    feed_name: str
    author_name: str
    feed_last_updated: datetime
    entries: list[Entry]


def get_feed_content() -> FeedContent:
    """Parse `feed.atom`, `public/`, and Git metadata to generate `FeedContent`."""
    # STUB: Parse content
    return FeedContent(
        base_url=BASE_URL,
        feed_name=FEED_NAME,
        author_name=AUTHOR_NAME,
        feed_last_updated=datetime.now(),
        entries=[],
    )


def generate_file_contents(feed: FeedContent) -> str:
    """Generates the entire `feed.atom` file as a string."""

    final_str = ""
    final_str += f"""<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns="http://www.w3.org/2005/Atom">
  <title>{feed.feed_name}</title>
  <id>{feed.base_url}</id>
  <link rel="alternate" href="{feed.base_url}"/>
  <link rel="self" href="{feed.base_url+"/feed.atom"}"/>
  <updated>{feed.feed_last_updated.isoformat()}</updated>
  <author>
    <name>{AUTHOR_NAME}</name>
  </author>
  <generator version="{FEED_GENERATOR_VERSION}">
    Luis Victoria's Atom Feed Generator
  </generator>
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
