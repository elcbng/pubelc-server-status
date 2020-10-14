config = {
  "server_url":".",
  "default_info":"Server disconnected"
}

function ready(fn) {
    if (document.readyState != 'loading'){
      fn();
    } else {
      document.addEventListener('DOMContentLoaded', fn);
    }
}
function ajax(type, url, success, error) {
  var request = new XMLHttpRequest()
  request.open(type, config["server_url"]+url, true)
  request.onload = function(){
    success(this.status,this.response)
  }
  request.onerror = function(){
    error()
  }
  request.send()
}

ready(function(){
  console.log("start init")
  ajax('GET', '/status.json', function(state, res){
    console.log('success',state,res)
  }, function(){
    console.log('error')
  })
})