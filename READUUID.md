来自于 [iOS 如何获取 Mach-O 的 UUID](https://www.jianshu.com/p/9201d5e34eb6)  

```
#import "LDAPMUUIDTool.h"
#import <mach-o/ldsyms.h>

#include <limits.h>
#include <mach-o/dyld.h>
#include <mach-o/nlist.h>
#include <string.h>

static NSMutableArray *_UUIDRecordArray;

@implementation LDAPMUUIDTool

+ (NSDictionary *)getUUIDDictionary {
    
    int imageCount = (int)_dyld_image_count();
    
    for(int iImg = 0; iImg < imageCount; iImg++) {
        
        JYGetBinaryImage(iImg);
    }
    
    return @{};
}
// 获取 Load Command, 会根据 header 的 magic 来判断是 64 位 还是 32 位
static uintptr_t firstCmdAfterHeader(const struct mach_header* const header) {
    switch(header->magic) {
        case MH_MAGIC:
        case MH_CIGAM:
            return (uintptr_t)(header + 1);
        case MH_MAGIC_64:
        case MH_CIGAM_64:
            return (uintptr_t)(((struct mach_header_64*)header) + 1);
        default:
            return 0;
    }
}

bool JYGetBinaryImage(int index) {
    const struct mach_header* header = _dyld_get_image_header((unsigned)index);
    if(header == NULL) {
        return false;
    }
    
    uintptr_t cmdPtr = firstCmdAfterHeader(header);
    if(cmdPtr == 0) {
        return false;
    }
    
    uint8_t* uuid = NULL;
    
    for(uint32_t iCmd = 0; iCmd < header->ncmds; iCmd++)
    {
        struct load_command* loadCmd = (struct load_command*)cmdPtr;
        
        if (loadCmd->cmd == LC_UUID) {
            struct uuid_command* uuidCmd = (struct uuid_command*)cmdPtr;
            uuid = uuidCmd->uuid;
            break;
        }
        cmdPtr += loadCmd->cmdsize;
    }
    
    const char* path = _dyld_get_image_name((unsigned)index);
    NSString *imagePath = [NSString stringWithUTF8String:path];
    NSArray *array = [imagePath componentsSeparatedByString:@"/"];
    NSString *imageName = array[array.count - 1];
    
    NSLog(@"buffer->name:%@",imageName);
    
    const char* result = nil;
    
    if(uuid != NULL)
    {
        result = uuidBytesToString(uuid);
        NSString *lduuid = [NSString stringWithUTF8String:result];
        NSLog(@"buffer->uuid:%@",lduuid);
    }
    
    return true;
}
static const char* uuidBytesToString(const uint8_t* uuidBytes) {
    CFUUIDRef uuidRef = CFUUIDCreateFromUUIDBytes(NULL, *((CFUUIDBytes*)uuidBytes));
    NSString* str = (__bridge_transfer NSString*)CFUUIDCreateString(NULL, uuidRef);
    CFRelease(uuidRef);
    
    return cString(str);
}
const char* cString(NSString* str) {
    return str == NULL ? NULL : strdup(str.UTF8String);
}


@end
```
