const cluster = require('cluster');
const numCPUs = require('os').cpus().length;
const WebSocket = require('ws');
const args = process.argv;
var url = `ws://${args[2]}`;

var count = 0;
var totalUser = args[3]/numCPUs
var spawnUser = 40
var spawnTime = 60 // ms


var inloop = totalUser / spawnUser

if(cluster.isMaster) {
  for (var i = 0; i < numCPUs; i++) {
    cluster.fork();
  }
} else {
    console.log("MAIN")
    // main();
}

function main() {
    console.log(`Total user ${totalUser} | Rate ${(spawnUser*1000/spawnTime)} user/sec | ${spawnTime} user / ${spawnTime} ms`)
    var mainInterval = setInterval(
        function() {
            if(++count > inloop){
                console.log(`Total user ${totalUser} | Rate ${(spawnUser*1000/spawnTime)} user/sec Finihed`)
                clearInterval(mainInterval)
                return
            }
            spawnUsers(count);
        }, spawnTime
    )
}

function spawnUsers(i) {
    var j = i + spawnUser
    setInterval(
        function() {
            if(++i > j){
                return
            }

            var c = new WebSocket(url);
            c.onopen = function(){
                var mes = new Object();
                mes.content = "Hello From : " + i
                setInterval(function(){c.send(JSON.stringify(mes))}, getRandom() )
            }
        }, 15
    );
}

function getRandom(){
    return Math.floor(Math.random() * 80000) + 20000;
}