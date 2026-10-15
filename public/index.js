import { registerHelloWorldComponent } from "./components/hello-world.js";
import { registerFooterComponent } from "./components/footer.js";

const app = () => {
    registerHelloWorldComponent();
    registerFooterComponent();
}

document.addEventListener('DOMContentLoaded', app);
