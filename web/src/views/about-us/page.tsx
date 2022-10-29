import { onMount, onCleanup } from "solid-js";

export default function AboutUsPage() {
  onMount(() => {
    document.title = "About Us";
  });
  onCleanup(() => {
    document.title = "Loading...";
  });
  return (
    <div>
      <h1>About Us</h1>
    </div>
  );
}
