"""feed.atom generator"""
# Using https://kevincox.ca/2022/05/06/rss-feed-best-practices/ as a guide

import argparse

# CONSTANTS
FEED_NAME: str = "LV"
HOMEPAGE_URL: str = "https://luis.vi"
AUTHOR_NAME: str = "Luis Victoria"

TEMPLATE: str = f"""<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns="http://www.w3.org/2005/Atom">
	<title>{FEED_NAME}</title>
	<id>{HOMEPAGE_URL}</id>
	<link rel="alternate" href="{HOMEPAGE_URL}"/>
	<link rel="self" href="{HOMEPAGE_URL+"/feed.atom"}"/>
	<updated>{{LAST_UPDATE_TIME in RFC3339 format}}</updated>
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
    prarser = argparse.ArgumentParser(
        description="`feed.atom` generator."
    )

    return parser.parse_args()


def main() -> None:
    args = parse_args()
    print("Hello world")


if __name__ == "__main__":
    main()
