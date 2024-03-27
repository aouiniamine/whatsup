import io from "socket.io-client"
import { HOST, SERVER } from "./envirement"

export const socket = new WebSocket(`${SERVER}/ws`)

// let ws = new WebSocket('ws://localhost:4000');