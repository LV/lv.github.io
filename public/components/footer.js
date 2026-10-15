class FooterComponent extends HTMLElement {
    connectedCallback() {
        this.innerHTML = `
            <footer>
                <hr />
                <div class="contact">
                    <a href="https://github.com/lv">GitHub</a>
                    <a href="https://www.linkedin.com/in/luisvictoria">LinkedIn</a>
                    <a href="https://www.youtube.com/@luisvictoria">YouTube</a>
                    <a href="https://www.twitch.tv/luisvictoriajr">Twitch</a>
                    <a href="https://lichess.org/@/lavp">Lichess</a>
                    <a href="${import.meta.resolve('../feed.atom')}">RSS</a>
                </div>
                <br>
                <p>&copy; 2016-2026 Luis Victoria</p>
            </footer>
        `;
    }
}

export const registerFooterComponent = () => customElements.define('lv-footer', FooterComponent);
