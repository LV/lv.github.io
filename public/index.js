import { registerHelloWorldComponent } from "./components/hello-world/hello-world.js";
import { registerFooterComponent } from "./components/footer/footer.js";

const app = () => {
    registerHelloWorldComponent();
    registerFooterComponent();
}

document.addEventListener('DOMContentLoaded', app);
