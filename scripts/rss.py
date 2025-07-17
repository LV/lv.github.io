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


@dataclass
class Entry:
    """An RSS entry."""
    title: str
    url: str
    permalink_id: str
    published: datetime
    last_updated: datetime
    html_content: str


def generate_file(feed_timestamp: datetime) -> str:
    """Generates the entire `feed.atom` file as a string."""
    # TODO: Generate entries
    header: str = f"""<?xml version="1.0" encoding="UTF-8"?>
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

    # Check for entries in `SITE_DIR`, and generate them accordingly

    footer: str = "</feed>"


def main() -> None:
    """Script entrypoint."""
    # Check if `feed.atom` exists
    # Check against current `feed.atom` and site-tree cache
    # Update `feed.atom` file accordingly


if __name__ == "__main__":
    main()
