import { blogPostsComponent } from "../../components/blog-posts.js";

const app = () => {
  blogPostsComponent();
}

// Ensure that HTML is fully loaded before registering components by using `async`
document.addEventListener('DOMContentLoaded', app);
