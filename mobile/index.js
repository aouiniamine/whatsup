import { AppRegistry, } from "react-native";
import { registerRootComponent } from "expo";
import App from "./App";

AppRegistry.registerComponent("whatsup", () => App);
registerRootComponent(App);