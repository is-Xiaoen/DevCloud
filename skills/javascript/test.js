// function sum (x , y) {
//     return x + y
// }

// console.log(sum(1, 10)) 

// var person = {name: '小明', age: 23}

// console.log(person)

// person.greet = function() {
//     console.log(`hello, my name is ${this.name}, age ${this.age}`)
// }
// person.greet()


// var name = '李四'
// var age = 23

// function greet () {
//     console.log(`hello, my name is ${name}, age ${age}`)
// }

// greet()


// fn = function (x , y) {
//     return x + y
// }
// console.log(fn(1,2)) 


// fn = (x, y) => {
//     return x + y
// }
// console.log(fn(1,2)) 

// fn = (x, y) => { return x + y}
// console.log(fn(1,2)) 

// fn = (x, y) => x + y
// console.log(fn(1,2)) 

// pow = x => x * x

// console.log(pow(100))

// pkg
import {firstName, sum as my_sum} from './tool.mjs'
console.log(my_sum(1,10)) 
console.log(firstName)

// pkg.
// import { MYAPP } from './app.mjs'
// console.log(MYAPP.version)



// pkg 
// import { default as MYAPP } from './app.mjs'
import appPKg from './app.mjs'
console.log(appPKg.name) 

// for item := rang a {}
var a = ['A', 'B', 'C'];
for (var item of a) {
    // fn(item)
    console.log(item)
}

// for k,v := range o {}
var o = {
    name: 'Jack',
    age: 20,
    city: 'Beijing'
};
for (var item of Object.keys(o)) {
    console.log(item, o[item])
}

// forEach
a.forEach((item) => {
    console.log(item)
})