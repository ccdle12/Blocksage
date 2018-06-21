import axios from 'axios';

// https://github.com/axios/axios
export default {
  getTx(txId) {
    return axios.get(`${(process.env.API_BASE_URL)}/txs/${txId}`)
  },
}