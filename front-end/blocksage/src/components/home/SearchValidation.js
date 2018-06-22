export default {
  validate: function(search_value) {
    let acceptable = false
    let value = 'None'
  
    switch(search_value.length) {
      case 64:
        acceptable = true
        value = 'tx'
        break
    }
  
    return { valid: acceptable, type: value }
  }
}