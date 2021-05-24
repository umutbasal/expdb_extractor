const data = require('./out.json')
let res = data

// res = res.filter(a=>{
//     if (a.cve?.length>0){
//         return a
//     }
// })

// 1827 total
// 330 with cves

console.log(res.length)
function getRandomInt(max) {
    return Math.floor(Math.random() * max);
}

function run(){
    let rand = getRandomInt(res.length)
    return (res[rand])
}

console.log(run())