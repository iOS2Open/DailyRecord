# infer 学习

## 参考文章
[Getting started with Infer](https://fbinfer.com/docs/getting-started.html#install-from-source)  
[分析 APP 或者其他项目](https://infer.liaohuqiu.net/docs/analyzing-apps-or-projects.html)  
[项目代码监控优化 -- 利用infer进行代码质量检测](https://www.jianshu.com/p/bcddae9767b9)  
[iOS-App接入infer静态分析扫描工具](https://www.jianshu.com/p/52f61498e2b2)    
[Infer bug types](https://infer.liaohuqiu.net/docs/infer-bug-types.html)

## 一、安装
> brew install infer

## 二、单个文件
> infer -- clang -c Hello.m

## 三、项目
> infer --keep-going -- xcodebuild -project Dev.xcodeproj  -scheme Dev

## 四、详细介绍
from [iOS-App接入infer静态分析扫描工具](https://www.jianshu.com/p/52f61498e2b2)

下面简单描述一下通过infer扫描出来的问题：  
### 4.1. DIRECT_ATOMIC_PROPERTY_ACCESS。

在代码中使用了使用了一个atomic的成员变量，infer建议我们将atomic修改为nonatomic。由于OC中，属性会被默认设置为atomic属性，我们需要显示将属性声明为nonatomic。关于atomic与nonatomic的区别可以参见文章https://my.oschina.net/linxiaoxi1993/blog/381332  
该警告占了大概800个。在代码中主动设置成员变量的nonatomic属性，即可去除警告  
### 4.2 ASSIGN_POINTER_WARNING  
由于在mrc时代，没有weak指针，所以一些view的属性声明是_、unsafe__unretain__的形式，在arc中，这个属性被判断为assign，需要将其修改为weak或者strong

### 4.3 NULL_DEREFERENCE  
空指针的情况。根据具体代码的不同，出现空指针的情况也有所不同。
- 1.传参为0的情况下。例如代码中，在调用showAlertViewA()时，将tag传参为0,infer检测此处传0，判断为一个NULL空指针，所以爆出警告。这里可以理解为误报，不会出现问题。  
- 2．通过malloc,calloc,realloc等函数申请内存，当内存不足时，有可能会在该函数中返回NULL，如果没有做NULL的判断，则警告  
- 3．在创建NSArray或者NSDictionary时，传入的参数有可能会nil。由于NSArray与NSDictionary不接受空指针，所以在对其addObject或者setObject:forKey: 时需要进行判断一下是否为nil。

### 4.4 IVAR_NOT_NULL_CHECKED
在代码中调用block，运行代码时，没有做判空处理。即需要改动为，

if(block){block()}

### 4.5 BAD_POINTER_COMPARISON
Implicitly checking whether NSNumber pointer is nil。没有判断一个NSNumber类型的对象是不是空？此处应该是误报。

### 4.6 TAINTED_VALUE_REACHING_SENSITIVE_FUNCTION
代码中使用了cookie的value。可以理解为误报

### 4.7 PARAMETER_NOT_NULL_CHECKED
传参时没有判断是否为null，加一次判断就可以了

### 4.8 STRONG_DELEGATE_WARNING
将一个delegate属性设置为strong的类型。

### 4.9  PREMATURE_NIL_TERMINATION_ARGUMENT
没有判断是否为空

### 4.10 REGISTERED_OBSERVER_BEING_DEALLOCATED
创建一个对象后，监听了某些通知，但是没有在dealloc中释放该通知。项目中出现这种问题的类，基本都是单例，不会被销毁。

### 4.11 MEMORY_LEAK
内存泄露。项目代码全面启动了ARC进行内存管理，在OC层没有扫描出内存泄露。目前扫描出的内存泄露问题都是使用了malloc或者ralloc等c语言内存申请函数，在函数提前return前没有及时free。



