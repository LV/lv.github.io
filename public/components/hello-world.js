/**
 * Defines and registers a custom HTML element `<lv-hello-world>`.
 * @extends {HTMLElement}
 */
class HelloWorldComponent extends HTMLElement {
    /** @returns {void} */
    connectedCallback() {
        this.textContent = 'hello world!';
    }
}

export const helloWorldComponent = () => customElements.define("lv-hello-world", HelloWorldComponent);
