var cluster = require('cluster');

if(cluster.isMaster) {
  for (var i = 0; i < 4; i++) {
    cluster.fork();
  }
} else {
  main();
}

function main() {
  setInterval(() => {
    console.log("MAIN")
  }, 1000)
}