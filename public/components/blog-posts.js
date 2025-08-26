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
        const xmlDoc = new DOMParser().parseFromString(xmlText, "text/xml");

        const parserError = xmlDoc.querySelector('parsererror');
        if (parserError) {
          throw new Error('Invalid XML');
        }

        // Iterate across all `<entry>` tags in the feed
        const entries = xmlDoc.querySelectorAll('entry');
        console.log(`Found ${entries.length} entries`);
      })
      .catch(error => {
        this.textContent = `Error: ${error.message}`;
      });
  }
}

export const blogPostsComponent = () => customElements.define('lv-blog-posts', LatestPosts);
