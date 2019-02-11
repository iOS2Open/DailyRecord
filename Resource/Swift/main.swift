//
//  main.swift
//  ASwiftDev
//
//  Created by ZhuHong on 2018/10/31.
//  Copyright © 2018年 CoderHG. All rights reserved.
//

import Foundation

// 函数中使用泛型
func num<T>(a:T) -> T {
    return a
}

// 使用
let var_int = num(a: 1)
let var_double = num(a:1.1)
// 打印
print(var_int)
print(var_double)

// 协议中使用泛形型
protocol Animal {
    associatedtype T
    func run() -> T
    func music() -> T
}

class Person: Animal {
    
    func run() -> Person {
        print("Person - run")
        return self
    }
    
    func music() -> Person {
        print("Person - music")
        return self
    }
}

class Pig: Animal {
    
    func run() -> Pig {
        print("Pig - run")
        return self
    }
    
    func music() -> Pig {
        print("Pig - music")
        return self
    }
}

let p = Person()
p.run().music()

let pig = Pig()
pig.run().music()


// 指定泛型的类型
class BigPig: Pig {
}

func funDefine<T>(a:T) where T:Pig {
}

// 直接报错
//funDefine(a: Person())
// 正确
funDefine(a: BigPig())

// 闭包的简单使用
func funcBody(a:Int, b:Int, block:(Int, Int)->(Int)) -> Int {
    return block(a, b)
}

// 加法运算
func sum(a:Int, b:Int) -> Int {
    let c = a + b
    return c
}
// 减法运算
//func sub(a:Int, b:Int) -> Int {
//    let c = a - b
//    return c
//}
// 或者
let sub:(Int, Int) -> Int = {
    (a, b) -> Int in
    return a - b
}

// 加法
var value = funcBody(a: 5, b: 2, block: sum)
print(value)
// 结果为 7

// 减法
value = funcBody(a: 5, b: 2, block: sub)
print(value)
// 结果为 3

// 乘法
value = funcBody(a: 5, b: 2, block: { (aa, bb) -> (Int) in
    return aa * bb
})
print(value)
// 结果为 10

// 除法
value = funcBody(a: 15, b: 3) { (aa, bb) -> (Int) in
    return aa/bb
}
print(value)
// 结果为 5


// 逃逸闭包
func taoyi(block:@escaping ()->Void) -> Void {
    // 这里与 OC 中的 Block 不相同
    // 在 OC 中, Block 中使用对象,会在 Block 对象中自动生成一个对应对应的引用(强引用/若引用)
    // 但是在 Swift 中却不能
    // 故一旦在 Swift 中,闭包中使用闭包的情况下, 想要延长其被引用的闭包, 那么需要将此闭包声明成逃逸闭包, 即添加关键词 @escaping
    let q = DispatchQueue(label: "label")
    let time = DispatchTime.now() + DispatchTimeInterval.seconds(2)
    q.asyncAfter(deadline: time) {
        block()
    }
    
    // 那么问题来了, 如何确定需要逃逸闭包:
    // 1. Xcode 会自动提示
    // 2. 像上面的例子, asyncAfter函数的中的闭包已经加有 @escaping 关键字
}
//调用闭包函数
taoyi {
    print("逃逸闭包的使用")
}

// 定义一个 Class
class WeakPerson {
    // 懒加载
    lazy var name:String = {
        return "CoderHG"
    }()
    
    var block:((Int)->())?
    func test() {
        // weak var weakSelf = self
        // unowned let weakSelf = self
        block = {
            /** [weak self] */[unowned self] (a) in
            print(self, a)
            // print(weakSelf ?? "")
        }
        block!(9)
    }
    
    deinit {
        print("释放了", self)
    }
}

var wp = WeakPerson()
wp.test()

wp = WeakPerson()
// **unowned** 相当于 OC 中的 **__unsafe_unretained**。

// 模仿 Masory
class CalculatorMaker {
    var result:Int = 0
    func sum(a:Int) -> CalculatorMaker {
        result += a;
        return self
    }
    func sub(a:Int) -> CalculatorMaker {
        result -= a;
        return self
    }
    
}
class Calculator {
    static func start(block:(CalculatorMaker) -> ()) -> Int {
        let cal = CalculatorMaker()
        block(cal)
        return cal.result
    }
}

let s = Calculator.start { (maker) in
    maker.sum(a: 2).sub(a: 3).sum(a: 5)
}
print(s)
// 结果: 4

print("Hello, World!")


// 定义一个结构体
struct HGStruct {
    var name:String?
    var des:String?
    
    func desFunc() -> Void {
        print(name! + "_" + des!)
    }
}

// 可以这样使用：
// 无参构造函数
var st = HGStruct()
// 逐一构造函数
st = HGStruct(name: "HG", des: "Good")

// 调用结构体函数
st.desFunc()

