const rpcOptions = require('../utils/rpc-options')

describe('rpcOptions', () => {
    it('should return a request options JSON object', () => {
        const options = rpcOptions.createOptions('getnetworkinfo')
        expect(options).toBeDefined()
    })

    it('should return the method call getnetworkinfo ', () => {
        const options = rpcOptions.createOptions('getnetworkinfo')
        expect(options.body).toEqual(JSON.stringify({"method": "getnetworkinfo"}))
    })
})
