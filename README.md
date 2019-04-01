# DailyRecord
 日常笔录
 
 ```
 - (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view, typically from a nib.
    
    
    NSMutableArray *arrM = [NSMutableArray array];
    [arrM addObject:@"4"];
    [arrM addObject:@"6"];
    [arrM addObject:@"8"];
    [arrM addObject:@"12"];
    [arrM addObject:@"15"];
    
    
    NSMutableArray *arrM1 = [NSMutableArray array];
    [arrM1 addObject:@"1"];
    [arrM1 addObject:@"3"];
    [arrM1 addObject:@"5"];
    [arrM1 addObject:@"7"];
    [arrM1 addObject:@"11"];
    [arrM1 addObject:@"18"];
    
    NSInteger tempInteger = arrM1.count-1;
    for (NSInteger j=(arrM.count-1); j>=0; j--) {
        NSString *number  = arrM[j]; // 新的
        for (NSInteger i=tempInteger; i>=0; i--) {
            NSString *lastIMMessage = arrM1[i];
            if (j==0) {
                
                if (number.integerValue < lastIMMessage.integerValue) {
                    [arrM insertObject:lastIMMessage atIndex:(j+1)];
                    tempInteger = i-1;
                }
                
                if (tempInteger >= 0) {
                    NSArray *arr = [arrM1 subarrayWithRange:NSMakeRange(0, i)];
                    [arrM insertObjects:arr atIndexes: [NSIndexSet indexSetWithIndexesInRange:NSMakeRange(0, i)]];
                }
                
                break;
            } else if (number.integerValue < lastIMMessage.integerValue) {
                [arrM insertObject:lastIMMessage atIndex:(j+1)];
                tempInteger = i-1;
                break;
            }
        }
    }
    
    
    NSLog(@"%@", arrM);
    
    
}
 ```

ok 

 ```
 // 把 arrMessages d放进 arrM 中
- (void)arrMessages:(NSArray*)arrMessages tempArrM:(NSMutableArray*)tempArrM {
    NSInteger tempInteger = arrMessages.count-1;
    for (NSInteger j=(tempArrM.count-1); j>=0; j--) {
        Message *imMessage  = tempArrM[j]; // 新的
        for (NSInteger i=tempInteger; i>=0; i--) {
            Message *lastIMMessage = arrMessages[i];
            if (j==0) {
                
                if (imMessage.timestamp < lastIMMessage.timestamp) {
                    [tempArrM insertObject:lastIMMessage atIndex:(j+1)];
                    tempInteger = i-1;
                }
                
                if (tempInteger >= 0) {
                    NSArray *arr = [arrMessages subarrayWithRange:NSMakeRange(0, i)];
                    [tempArrM insertObjects:arr atIndexes: [NSIndexSet indexSetWithIndexesInRange:NSMakeRange(0, i)]];
                }
                
                break;
            } else if (imMessage.timestamp < lastIMMessage.timestamp) {
                [tempArrM insertObject:lastIMMessage atIndex:(j+1)];
                tempInteger = i-1;
                break;
            }
        }
    }
}
 ```

重连机制：  
```
- (void)reConnectServer {
    
    // 静态变量 记录重连的次数
    static NSInteger connectAcount = 0;
    // 静态变量 记录上一次的时间
    static CFTimeInterval sLastTimeInterval = 0;
    
    // 获取当前时间
    CFTimeInterval timeInterval = CACurrentMediaTime();
    
    
    @synchronized (self) {
        // 记录数增加
        connectAcount ++;
        
        // 超过3分钟, 强制记录重置
        if ((timeInterval - sLastTimeInterval) > 3*60) {
            // 重新l记录时间节点
            sLastTimeInterval = timeInterval;
            // 次数重置
            connectAcount = 0;
        }
    }
    
    // 时间间隙
    double margin = 0.0;
    
    /**  规则 connectAcount
     0, 没有间隔
     在 0  ~ 10 次之内, 间隔 0.1 秒
     在 10 ~ 20 次之内, 间隔 0.5 秒
     在 20 ~ 30 次之内, 间隔 1.0 秒
     在 30 ~ 40 次之内, 间隔 1.5 秒
     其余 2.0 秒
     */
    if (connectAcount == 0) {
         margin = 0.0;
    } else if (connectAcount < 11) {
        margin = 0.1;
    } else if (connectAcount < 21) {
        margin = 0.5;
    } else if (connectAcount < 31) {
        margin = 1.0;
    } else if (connectAcount < 41) {
        margin = 1.5;
    } else {
        margin = 2.0;
    }
    
    if (margin == 0.0) {
        [self reConnectServer];
    } else {
        dispatch_after(dispatch_time(DISPATCH_TIME_NOW, (int64_t)(margin * NSEC_PER_SEC)), dispatch_get_main_queue(), ^{
            [self reConnectServer];
        });
    }
}
```
