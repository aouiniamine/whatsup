import axios from 'axios'
import { SERVER } from './envirement'
export const connect = (credential) =>
    axios.post(`${SERVER}/auth/connect`, {credential})
    .then(res => {
        if (res.data?.email){
            return res.data
        } else throw res.data
    })

export const register = (credential) =>
    axios.post(`${SERVER}/auth/register`, credential)
    .then(res => {
        if (res.data?.email){
            return res.data
        } else throw res.data
    })

export const validate = (body) =>
    axios.post(`${SERVER}/auth/validate`, body)
    .then(res => res.data)