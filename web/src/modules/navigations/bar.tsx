import { A } from "@solidjs/router";

export default function NavigationBar() {
  return (
    <nav class="flex flex-row justify-between items-center px-4 py-2 w-screen h-[42px] fixed top-0 bg-zinc-50">
      <div class="cursor-default">
        <strong class="text-xl">📦</strong>
      </div>
      <div>
        <A href="/">Home</A> <A href="/about-us">About Us</A>
      </div>
    </nav>
  );
}
