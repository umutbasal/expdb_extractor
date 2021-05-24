const data = require('./out.json')
let res = data

function getRandomInt(max) {
    return Math.floor(Math.random() * max);
}

function run(){
    let rand = getRandomInt(res.length)
    return (res[rand])
}

console.log(run())