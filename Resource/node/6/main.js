// 输出全局变量 __filename 的值
console.log( __filename );

function printHello(){
   console.log( "Hello, World!");
}
// 两秒后执行以上函数
setTimeout(printHello, 2000);


// 输出全局变量 __dirname 的值
console.log( __dirname );

// 两秒后执行以上函数
setInterval(printHello, 2000);

console.error('12344')

// 获取执行路径
console.log(process.execPath);


// 平台信息
console.log(process.platform);
