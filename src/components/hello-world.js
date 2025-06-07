/**
 * Defines and registers a custom HTML element `<lv-hello-world>`.
 * @returns {void}
 */
class HelloWorldComponent extends HTMLElement {
    connectedCallback() {
        this.textContent = 'hello world!';
    }
}

export const helloWorldComponent = () => customElements.define("lv-hello-world", HelloWorldComponent);
