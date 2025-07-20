function testApiCall(success, failed) {
    var timeOut = Math.random() * 2;
    console.log('set timeout to: ' + timeOut + ' seconds.');
    setTimeout(function () {
        if (timeOut < 1) {
            console.log('call resolve()...');
            success('200 OK');
        }
        else {
            console.log('call reject()...');
            failed('timeout in ' + timeOut + ' seconds.');
        }
    }, timeOut * 1000);
}

// 通过回调实现的异步方案
// fna()
// fnb()
// fnc()
console.log('start ...')
testApiCall(
    // success callback
    (data) => {
        console.log(data)
    },
    // failed callback
    (error) => {
        console.log(error)
    }
)
console.log('end ...')

// fna(fnb_sucess(fnc_success(fnd_success)), fnb_error(fnc_error(fnd_error)))

// promise 封装技术, 他让异步成为一种标准, 使用promise的标准方式来处理异常
// promise 像一个接口，他约束了 excutor对象的 实现方式
var p1 = new Promise(testApiCall)
// p1.then((data) => {
//     console.log(data)
//     //  p2.then().catch().finally()
// }).catch((error) => {
//     console.log(error)
// }).finally(() => {
//     console.log('finnal')
// })


// 异步等待
// 一同步编程的方式，来实现异步代码
var async_await = async () => {
    try {
        const resp = await p1
        console.log(resp)

        // const res2 = await p2
        // console.log(resp2)

        // const res3 = await p3
        // console.log(resp2)
    } catch (error) {
        console.log(error)
    }
}

async_await()




