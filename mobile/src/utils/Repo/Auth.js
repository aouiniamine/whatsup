import axios from 'axios'
import { SERVER } from './envirement'
export const connect = (credentail) =>
    axios.post(`${SERVER}/auth/connect`, {credentail}, {
        headers: {
            "Content-Type": "application/json"
        }
    })