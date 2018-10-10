import axios from 'axios'

// https://github.com/axios/axios
export default {
  getTX(txID) {
    return axios.get(`${(process.env.API_BASE_URL)}/txs/${txID}`)
  },
}