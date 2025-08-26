// Refined from https://plainvanillaweb.com/blog/components/blog-latest-posts.js

import { html } from '../lib/html.js';

/** @type {number} */
const ENTRIES_TO_SHOW = 6;

class LatestPosts extends HTMLElement {
  // `connectedCallback` immediately loads; doesn't need explicit call to a function
  connectedCallback() {
    // Fallback text to display while content loads
    this.textContent = "Loading...";

    fetch('/feed.atom')
      .then(response => response.text())
      .then(xmlText => {
        // Convert the XML file string into queryable DOM
        /** @type {Document} */
        const xmlDoc = new DOMParser().parseFromString(xmlText, "text/xml");

        /** @type {Element | null} */
        const parserError = xmlDoc.querySelector('parsererror');
        if (parserError) {
          throw new Error('Invalid XML');
        }

        // Iterate across all `<entry>` tags in the feed
        /** @type {NodeListOf<Element>} */
        const entries = xmlDoc.querySelectorAll('entry');
        console.log(`Found ${entries.length} entries`);

        const feedItems = Array.from(entries)
          .slice(0, ENTRIES_TO_SHOW)
          .map(entry => ({
            title: entry.querySelector('title')?.textContent,
            link: entry.querySelector('link')?.getAttribute('href'),
            published: entry.querySelector('published')?.textContent,
            summary: entry.querySelector('summary')?.textContent
          }))
          .filter(item => item.title && item.link);

        console.log('Feed items:', feedItems);
      })
      .catch(error => {
        this.textContent = `Error: ${error.message}`;
      });
  }
}

export const blogPostsComponent = () => customElements.define('lv-blog-posts', LatestPosts);
