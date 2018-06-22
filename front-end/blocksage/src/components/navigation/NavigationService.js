export default {
    navigate: function(router, searchRequest, searchParams) {
        if (searchRequest.valid && searchRequest.type == 'tx') 
            router.push({ name: searchRequest.type, params: { txHash: searchParams } })

        return false
    },
  }