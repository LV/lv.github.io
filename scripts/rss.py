"""feed.atom generator"""
# Using https://kevincox.ca/2022/05/06/rss-feed-best-practices/ as a guide

import argparse
from dataclasses import dataclass
from datetime import datetime

# CONSTANTS
FEED_NAME: str = "Luis Victoria's Blog"
HOMEPAGE_URL: str = "https://luis.vi"
AUTHOR_NAME: str = "Luis Victoria"


@dataclass
class Entry:
    title: str
    url: str
    permalink_id: str
    published: datetime
    last_updated: datetime
    html_content: str


def generate_file(base_url: str, feed_timestamp: datetime) -> str:
    return f"""<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns="http://www.w3.org/2005/Atom">
  <title>{FEED_NAME}</title>
  <id>{base_url}</id>
  <link rel="alternate" href="{base_url}"/>
  <link rel="self" href="{base_url+"/feed.atom"}"/>
  <updated>{feed_timestamp.isoformat()}</updated>
  <author>
    <name>{AUTHOR_NAME}</name>
  </author>
  <entry>
    <title>{{ENTRY.TITLE}}</title>
    <link rel="alternate" type="text/html" href="{{ENTRY.HTML_URL}}"/>
    <id>{{ENTRY.PERMALINK}}</id>
    <published>{{ENTRY.FIRST_POST_TIME in RFC3339 format}}</published>
    <updated>{{ENTRY.LAST_UPDATE_TIME in RFC3339 format}}</updated>
    <content type="html">{{ENTRY.HTML}}</content>
  </entry>
</feed>
"""


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        prog="RSS Generator",
        description="Generate, edit, and modify the `feed.atom` for RSS support",
    )
    parser.add_argument(
        "-u",
        "--update-feed-timestamp",
        help="Update the feed's last modified timestamp to now",
    )
    # Feature request: Add blog post; should create a new entry in the RSS

    return parser.parse_args()


def main() -> None:
    args = parse_args()
    # Check if `feed.atom` file exists
    print("Hello world")


if __name__ == "__main__":
    main()
