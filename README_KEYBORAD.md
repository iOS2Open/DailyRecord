[关于IOS键盘的一些调用心得与坑](https://www.jianshu.com/p/e7fe06c739e6)  

```
-(BOOL)textField:(UITextField *)textField shouldChangeCharactersInRange:(NSRange)range replacementString:(NSString *)string{      if ([textField.text isEqualToString:@""]) { //这个时候得到的还是文本框未改变时的字符串          if ([string length] < 1 ) {              //如果用户点击的删除按键，那么string是空的，没有长度          }           if ([string length]>0) {                int result = [string characterAtIndex:0];                 if (result == 10) {                          //用户点击了回车                 }               i f (result == 32) {                  //用户点击了空格               } }}return YES;}

```
