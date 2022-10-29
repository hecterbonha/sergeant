import { lazy, Suspense } from "solid-js";
import NavigationBar from "../modules/navigations/bar";
import Footer from "../modules/navigations/footer";
import AppRouter from "./router";

const WalletPopUp = lazy(() => import("../modules/wallets/pop-up"));
const SideNavigation = lazy(() => import("../modules/navigations/side"));

export default function Core() {
  return (
    <>
      <NavigationBar />
      <Suspense fallback={<></>}>
        <SideNavigation />
      </Suspense>
      <div class="flex flex-row justify-center items-center min-h-screen bg-zinc-100 dark:bg-zinc-900 text-zinc-900 dark:text-zinc-100">
        <AppRouter />
      </div>
      <Footer />
      <Suspense fallback={<></>}>
        <WalletPopUp />
      </Suspense>
    </>
  );
}
