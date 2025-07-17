"""`feed.atom` Generator"""
# IDEA: Use Git `master` metadata to add timestamps and stuff to blogs and entries
# Using https://kevincox.ca/2022/05/06/rss-feed-best-practices/ as a guide

from dataclasses import dataclass
from datetime import datetime
import sys

# CONSTANTS
AUTHOR_NAME: str = "Luis Victoria"
BASE_URL: str = "https://luis.vi"
FEED_NAME: str = "Luis Victoria's Blog"


@dataclass
class Entry:
    title: str
    url: str
    permalink_id: str
    published: datetime
    last_updated: datetime
    html_content: str


def generate_file(feed_timestamp: datetime, entries: list[Entry]) -> str:
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

    if entries:

    footer: str = "</feed>"


def main() -> None:
    # Parse site tree (`public`)
        # Check if `feed.atom` exists
    # Check against current `feed.atom` and site-tree cache
    # Update `feed.atom` file accordingly


if __name__ == "__main__":
    main()
