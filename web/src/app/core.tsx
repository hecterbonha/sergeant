import { lazy } from "solid-js";
import NavigationBar from "../modules/navigations/bar";
import AppRouter from "./router";

const WalletPopUp = lazy(() => import("../modules/wallets/pop-up"));

export default function Core() {
  return (
    <>
      <NavigationBar />
      <div class="flex flex-row justify-center items-center min-h-screen bg-zinc-900 text-zinc-50">
        <AppRouter />
      </div>
      <WalletPopUp />
    </>
  );
}
