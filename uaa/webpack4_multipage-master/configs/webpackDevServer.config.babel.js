import path from "path";
import { appPath } from "./env";

export default {
  contentBase: path.resolve(appPath, "pages/**/*.html"),
  publicPath: "/",
  watchContentBase: true,
  overlay: true,
  open: false,
  stats: "errors-only",
  hot: true,
  before (app, server) {
    server._watch(path.resolve(appPath, "pages/**/*.html"));
  }
}