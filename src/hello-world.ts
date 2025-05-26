class HelloWorldComponent extends HTMLElement {
    connectedCallback(): void {
        this.textContent = 'hello world!';
    }
}

customElements.define('x-hello-world', HelloWorldComponent);
