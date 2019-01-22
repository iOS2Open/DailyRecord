## 参考文章
[iOS App冷启动治理：来自美团外卖的实践](https://mp.weixin.qq.com/s/jN3jaNrvXczZoYIRCWZs7w)
[objc_cover](https://github.com/Reference2GreatGod/objc_cover)  
[Hades：移动端静态分析框架 & 还没有来得及看](https://tech.meituan.com/hades.html)  

在看的过程中，发现这两个地方很好用，也可以直接使用。
## 一、代码瘦身
![image.png](https://upload-images.jianshu.io/upload_images/1198135-2c0a7d0474dda380.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

很不错的。

## 二、冷启动开始&结束时间
![meituan.png](https://upload-images.jianshu.io/upload_images/1198135-67c07893445d62cf.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

代码如下：  
```objc
#import <sys/sysctl.h>
#import <mach/mach.h>

+ (BOOL)processInfoForPID:(int)pid procInfo:(struct kinfo_proc*)procInfo
{
    int cmd[4] = {CTL_KERN, KERN_PROC, KERN_PROC_PID, pid};
    size_t size = sizeof(*procInfo);
    return sysctl(cmd, sizeof(cmd)/sizeof(*cmd), procInfo, &size, NULL, 0) == 0;
}

+ (NSTimeInterval)processStartTime
{
    struct kinfo_proc kProcInfo;
    if ([self processInfoForPID:[[NSProcessInfo processInfo] processIdentifier] procInfo:&kProcInfo]) {
        return kProcInfo.kp_proc.p_un.__p_starttime.tv_sec * 1000.0 + kProcInfo.kp_proc.p_un.__p_starttime.tv_usec / 1000.0;
    } else {
        NSAssert(NO, @"无法取得进程的信息");
        return 0;
    }
}
```

