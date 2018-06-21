import Home from './components/home/Home.vue'
import Tx from './components/tx/Tx.vue'

export const routes = [
    { path: '', component: Home },
    //TODO: Create route guard to make sure there is a txHash parameter
    { path: '/tx/:txHash', component: Tx, name: 'tx' }
]