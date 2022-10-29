import { A } from "@solidjs/router";
import { For } from "solid-js";

export default function NavigationBar() {
  const navigationItems = [
    {
      title: "Home",
      url: "/",
    },
    {
      title: "Learn",
      url: "/learn",
    },
    {
      title: "Experiment",
      url: "/experiment",
    },
    {
      title: "About Us",
      url: "/about-us",
    },
  ];
  return (
    <nav class="flex flex-row justify-between items-center px-4 py-2 w-[calc(100vw-72px)] h-[64px] fixed top-0 bg-zinc-50 border-b-2 border-b-zinc-200">
      <div class="cursor-default">
        <strong class="text-xl">📦</strong>
      </div>
      <div class="flex flex-row gap-4">
        <For each={navigationItems} fallback={<div>Loading...</div>}>
          {(item) => (
            <A
              href={item.url}
              inactiveClass="hover:border-b-2 hover:border-b-zinc500"
              activeClass="border-b-2 border-b-zinc500"
              end
            >
              {item.title}
            </A>
          )}
        </For>
      </div>
    </nav>
  );
}
