const base64_auth = require('../../utils/base64-auth')
const debug = require('debug')('app:base64-auth-test')

describe('base64Auth', () =>
{
    it('should encode environment variables username:password to base64', () =>
    {
        const result = base64_auth.getEncodedAuth()
        expect(result).toEqual('Yml0Y29pbjpwYXNzd29yZA==')
    })
})