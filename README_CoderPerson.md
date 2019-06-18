```
#import "CoderPerson.h"
#import <Foundation/Foundation.h>

// 有一个问题想要请教,可能与当前文章没有太大的关联.
// 在女神(念茜)的博客(iOS安全攻防（二十二）：static和被裁的符号表)中提到
/**
 如果函数属性为 static ，那么编译时该函数符号就会被解析为local符号。
 在发布release程序时（用Xcode打包编译二进制）默认会strip裁掉这些函数符号，无疑给逆向者加大了工作难度。
 */

// 但是根据我的测试发现, 实际上并不是女神说的那样.在 iOS 项目中, 一个 C 函数即使是不添加 static 关键字, 在 release 的时候, 也没有具体的符号表.
// 这是因为现在的 Xcode 做了优化了么?

@implementation CoderPerson


// 发现, 即使是不加 static, 在 release 的时候这个函数符号也没有
//  static int coder_func() {
int coder_func() {
    int a = 10;
    int b = 20;
    int c = a+b;
    return c;
}

- (void)coder_method {
    int d = coder_func();
    NSLog(@"%d", d);
}

@end
```
