import { registerHelloWorldComponent } from "./components/hello-world/hello-world.js";
import { registerHeaderComponent } from "./components/header/header.js";
import { registerFooterComponent } from "./components/footer/footer.js";

const app = () => {
    registerHelloWorldComponent();
    registerHeaderComponent();
    registerFooterComponent();
}

document.addEventListener('DOMContentLoaded', app);
