import { createSignal, Show } from "solid-js";

export const [popUp, triggerPopUp] = createSignal<boolean>(false);

export default function WalletPopUp() {
  return (
    <Show when={popUp()}>
      <div class="fixed bottom-4 right-4 bg-zinc-50 p-4">
        <h1>Wallet Pop Up</h1>
        <button
          onClick={() => {
            triggerPopUp(false);
          }}
        >
          Close Me
        </button>
      </div>
    </Show>
  );
}
