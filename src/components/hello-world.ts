class HelloWorldComponent extends HTMLElement {
    connectedCallback(): void {
        this.textContent = 'hello world!';
    }
}

export const helloWorldComponent = () => customElements.define("lv-hello-world", HelloWorldComponent);
