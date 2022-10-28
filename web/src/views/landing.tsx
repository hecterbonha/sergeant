import { onMount, onCleanup } from "solid-js";
import ThemeSwitcher from "../modules/utils/theme-switcher";
import { triggerPopUp } from "../modules/wallets/pop-up";
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
      <h1>{navigator}</h1>
      <br />
      <ThemeSwitcher />
      <br />
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
