import { registerHelloWorldComponent } from "./components/hello-world/hello-world.js";

const app = () => {
    registerHelloWorldComponent();
}

document.addEventListener('DOMContentLoaded', app);
