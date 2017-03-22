const fs = require('fs')
const util = require('util')

var file = fs.readFileSync('./citylots.json');

console.time(' [x] Node.js unmarshal');
var data = JSON.parse(file);
console.timeEnd(' [x] Node.js unmarshal');

console.time(' [x] Node.js marshal');
var str = JSON.stringify(data)
console.timeEnd(' [x] Node.js marshal');

// console.log(util.inspect(data.features[0], {depth: 4, colors: true}))