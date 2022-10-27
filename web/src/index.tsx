/* @refresh reload */
import { render } from "solid-js/web";
import { Router } from "@solidjs/router";

import "./index.css";
import Core from "./app/core";

render(
  () => (
    <Router>
      <Core />
    </Router>
  ),
  document.getElementById("instead")
);
