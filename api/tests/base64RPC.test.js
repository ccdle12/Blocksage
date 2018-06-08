const base64RPC = require('../utils/base64RPC')

describe('base64RPC', () => {
    it('should encode environment variables username:password to base64', () => {
        const result = base64RPC.getRPCAuth()
        expect(result).toEqual('Yml0Y29pbjpwYXNzd29yZA==')
    })
})
