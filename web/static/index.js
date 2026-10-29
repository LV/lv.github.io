import { registerHelloWorldComponent } from "./components/hello-world/hello-world.js";
import { registerHeaderComponent } from "./components/header/header.js";

const app = () => {
    registerHelloWorldComponent();
    registerHeaderComponent();
}

document.addEventListener('DOMContentLoaded', app);
