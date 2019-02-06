## Cookie  
iOS 中的 Cookie 获取方式如下：  

```
NSHTTPCookieStorage *cookieStorage = [NSHTTPCookieStorage sharedHTTPCookieStorage];
NSArray *cookies = [cookieStorage cookies];
for (NSHTTPCookie *ck in cookies) {
    NSLog(@"%@", ck);
}

cookies = [NSHTTPCookie requestHeaderFieldsWithCookies:[cookieStorage cookies]];
NSLog(@"%@", cookies);
NSLog(@"cookieStorage = %@", cookieStorage);
```
