/** @type {{ id: string, label: string, href: string }[]} */
const NAV_LINKS = [
    { id: 'home', label: 'Home', href: import.meta.resolve('../../index.html') },
    { id: 'blog', label: 'Blog', href: import.meta.resolve('../../blog/index.html') },
];

/**
 * Defines and registers a custom HTML element `<lv-header>`.
 * The `current` attribute marks which nav link is the active page.
 * @extends {HTMLElement}
 */
class HeaderComponent extends HTMLElement {
    static get observedAttributes() {
        return ['current'];
    }

    /** @returns {void} */
    connectedCallback() {
        this.update();
    }

    /** @returns {void} */
    attributeChangedCallback() {
        this.update();
    }

    /** @returns {void} */
    update() {
        if (!this.isConnected) return;

        const current = this.getAttribute('current');
        const links = NAV_LINKS.map(({ id, label, href }) => id === current
            ? `<li><a href="#" aria-current="page">${label}</a></li>`
            : `<li><a href="${href}">${label}</a></li>`
        ).join('');

        this.innerHTML = `
            <header>
                <h1>Luis Victoria</h1>
                <p>Creating to Understand</p>

                <nav>
                    <ul>${links}</ul>
                </nav>
            </header>
        `;
    }
}

export const registerHeaderComponent = () => customElements.define('lv-header', HeaderComponent);
