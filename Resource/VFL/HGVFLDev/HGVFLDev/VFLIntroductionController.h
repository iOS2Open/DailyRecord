//
//  VFLIntroductionController.h
//  HGVFLDev
//
//  Created  by hong.zhu on 2019/3/30.
//  Copyright © 2019年 UCARINC. All rights reserved.
//  可视化格式语言

#import <UIKit/UIKit.h>

NS_ASSUME_NONNULL_BEGIN

@interface VFLIntroductionController : UIViewController

@end

NS_ASSUME_NONNULL_END

/** 语法说明
H:[cancelButton(72)]-12-[acceptButton(50)]
cancelButton宽72，acceptButton宽50，它们之间间距12

H:[wideView(>=60@700)]
wideView宽度大于等于60point，该约束条件优先级为700（优先级最大值为1000，优先级越高的约束条件越先被满足）

V:[redBox][yellowBox(==redBox)]
垂直方向上，先有一个redBox，其下方紧接一个高度等于redBox高度的yellowBox

H:|-10-[Find]-[FindNext]-[FindField(>=20)]-|
水平方向上，Find距离父view左边缘间隔10，之后是FindNext距离Find间隔默认宽度；再之后是宽度不小于20的FindField，它和FindNext以及父view右边边缘的间距都是默认宽度。（竖线“|”表示superview的边缘）。
 */

/** 使用方法
 使用VFL来创建约束数组
 +(NSArray *)constraintsWithVisualFormat:(NSString *)format options:(NSLayoutFormatOptions)opts metrics:(NSDictionary *)metrics views:(NSDictionary *)views;
 
 format：VFL语句
 opts：约束类型
 metrics：VFL语句中用到的具体数值
 views：VFL语句中用到的控件
 
 创建一个字典（内部包含VFL语句中用到的控件）的快捷宏定义
 NSDictionaryOfVariableBindings(...)
 */
