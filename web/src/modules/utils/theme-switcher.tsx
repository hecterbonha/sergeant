import { createEffect, createSignal, onMount } from "solid-js";

type Theme = "light" | "dark" | "system";

export default function ThemeSwitcher() {
  const mode = localStorage.getItem("theme") as Theme;
  const [theme, setTheme] = createSignal<Theme>(mode);

  createEffect(() => console.log("count =", theme()));

  return (
    <select
      onInput={(e) => {
        const selected = (e.target as HTMLInputElement).value as Theme;
        localStorage.setItem("theme", selected);
        setTheme(selected);
      }}
      class="bg-zinc-600 p-4"
      value={theme()}
    >
      <option value="light">L</option>
      <option value="dark">D</option>
      <option value="system">S</option>
    </select>
  );
}
