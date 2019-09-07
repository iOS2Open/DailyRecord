```

#import <Foundation/Foundation.h>

/// 代理定义
@protocol IMSocketManagerDelegate <NSObject>

@optional
/// 连接, 断开连接   .....  等回调

@end

NS_ASSUME_NONNULL_BEGIN

@interface IMSocketManager : NSObject

/** 代理 */
@property (nonatomic, weak) id<IMSocketManagerDelegate> delegate;

/** 初始化 socket并连接 */
- (void)connectWithParams:(NSDictionary*)arams;

/** 断开连接 */
- (void)disConnectSocketWithCompletion:(void (^)(BOOL succeed, NSError * _Nullable error))completion;

/** 调用 emit 发送 Socket 消息 */
- (void)sendDataWithDictionary:(NSDictionary *)dictMsg handle:(void (^)(NSArray *data))handle;

@end

NS_ASSUME_NONNULL_END

```


sutop

apt-get

whoami

pgadmin

psql iosmanager
