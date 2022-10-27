import { lazy } from "solid-js";
import { Routes, Route } from "@solidjs/router";

const Home = lazy(() => import("../views/landing"));
const AboutUs = lazy(() => import("../views/about-us"));

export default function AppRouter() {
  return (
    <Routes>
      <Route path="/" component={Home} />
      <Route path="/about-us" component={AboutUs} />
    </Routes>
  );
}
