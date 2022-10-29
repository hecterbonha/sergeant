import { lazy } from "solid-js";
import { Routes, Route } from "@solidjs/router";

const Home = lazy(() => import("../views/home/page"));
const Learn = lazy(() => import("../views/learn/page"));
const Experiment = lazy(() => import("../views/experiment/page"));
const AboutUs = lazy(() => import("../views/about-us/page"));

export default function AppRouter() {
  return (
    <Routes>
      <Route path="/" component={Home} />
      <Route path="/learn" component={Learn} />
      <Route path="/experiment" component={Experiment} />
      <Route path="/about-us" component={AboutUs} />
    </Routes>
  );
}
