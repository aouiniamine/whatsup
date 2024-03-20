import axios from 'axios'
import { SERVER } from './envirement'
export const connect = (credential) =>
    axios.post(`${SERVER}/auth/connect`, {credential})