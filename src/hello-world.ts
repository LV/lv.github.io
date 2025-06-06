class HelloWorldComponent extends HTMLElement {
    connectedCallback(): void {
        this.textContent = 'hello world!';
    }
}

customElements.define('lv-hello-world', HelloWorldComponent);
