"""`feed.atom` Generator"""
# IDEA: Use Git `master` metadata to add timestamps and stuff to blogs and entries
# Using https://kevincox.ca/2022/05/06/rss-feed-best-practices/ as a guide

from dataclasses import dataclass
from datetime import datetime
from pathlib import Path
import sys

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


def generate_file_contents(feed_timestamp: datetime) -> str:
    """Generates the entire `feed.atom` file as a string."""

    final_str = ""
    final_str += f"""<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns="http://www.w3.org/2005/Atom">
  <title>{FEED_NAME}</title>
  <id>{BASE_URL}</id>
  <link rel="alternate" href="{BASE_URL}"/>
  <link rel="self" href="{BASE_URL+"/feed.atom"}"/>
  <updated>{feed_timestamp.isoformat()}</updated>
  <author>
    <name>{AUTHOR_NAME}</name>
  </author>
"""

    # TODO: Generate entries
    # Check for entries in `SITE_DIR`, and generate them accordingly

    final_str += "</feed>"
    return final_str


def write_feed_file(content: str) -> None:
    """Writes the `feed.atom` file."""
    if FEED_FILE.is_file():
        FEED_FILE.unlink()

    FEED_FILE.write_text(content)


def main() -> None:
    """Script entrypoint."""
    should_write_file: bool = False

    if not FEED_FILE.is_file():
        should_write_file = True
        write_feed_file(generate_file_contents(datetime.now()))
        return

    # Check against current `feed.atom` and site-tree cache (aka check if the file needs updating)
    # Update `feed.atom` file accordingly

    print(generate_file_contents(datetime.now()))


if __name__ == "__main__":
    main()
