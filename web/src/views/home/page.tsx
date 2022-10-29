import { onMount, onCleanup } from "solid-js";

export default function HomePage() {
  onMount(() => {
    document.title = "Hello World";
  });
  onCleanup(() => {
    document.title = "Loading...";
  });
  const navigator = window.navigator.userAgent;
  return (
    <div>
      <p class="text-sm">{navigator}</p>
    </div>
  );
}
