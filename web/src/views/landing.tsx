import { onMount, onCleanup } from "solid-js";
import { triggerPopUp } from "../modules/wallets/pop-up";
export default function HomePage() {
  onMount(() => {
    document.title = "Hello World";
  });
  onCleanup(() => {
    document.title = "Loading...";
  });
  return (
    <div>
      <h1>Home</h1>
      <button
        onClick={() => {
          triggerPopUp(true);
        }}
      >
        Trigger Wallet
      </button>
    </div>
  );
}
