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
