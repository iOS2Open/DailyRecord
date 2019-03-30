//
//  VFLIntroductionController.m
//  HGVFLDev
//
//  Created  by hong.zhu on 2019/3/30.
//  Copyright © 2019年 UCARINC. All rights reserved.
//

#import "VFLIntroductionController.h"

@interface VFLIntroductionController ()

@end

@implementation VFLIntroductionController

- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view.
    
    [self horizontalLayout];
}

-(void)horizontalLayout{
    //1.添加两个控件
    UIView *blueView = [[UIView alloc] init];
    blueView.backgroundColor = [UIColor blueColor];
    
    blueView.translatesAutoresizingMaskIntoConstraints = NO;
    [self.view addSubview:blueView];
    
    UIView *redView = [[UIView alloc] init];
    redView.backgroundColor = [UIColor redColor];
    redView.translatesAutoresizingMaskIntoConstraints = NO;
    [self.view addSubview:redView];
    
    //2.添加约束
    //2.1水平方向的约束
    NSString *hVFL = @"H:|-30-[blueView]-30-[redView(==blueView)]-30-|";
    NSArray *hCons = [NSLayoutConstraint constraintsWithVisualFormat:hVFL options:NSLayoutFormatAlignAllBottom | NSLayoutFormatAlignAllTop metrics:nil views:@{@"blueView":blueView, @"redView":redView}];
    [self.view addConstraints:hCons];
    
    //2.2垂直方向的约束
    NSString *vVFL = @"V:[blueView(50)]-30-|";
    NSArray *vCons = [NSLayoutConstraint constraintsWithVisualFormat:vVFL options:0 metrics:nil views:@{@"blueView":blueView}];
    [self.view addConstraints:vCons];
}

-(void)verticalLayout{
    //1.添加两个控件
    UIView *blueView = [[UIView alloc] init];
    blueView.backgroundColor = [UIColor blueColor];
    blueView.translatesAutoresizingMaskIntoConstraints = NO;
    [self.view addSubview:blueView];
    
    UIView *redView = [[UIView alloc] init];
    redView.backgroundColor = [UIColor redColor];
    redView.translatesAutoresizingMaskIntoConstraints = NO;
    [self.view addSubview:redView];
    
    //2.添加约束
    //2.1水平方向的约束
    NSString *hVFL = @"H:|-30-[blueView]-30-|";
    NSArray *hCons = [NSLayoutConstraint constraintsWithVisualFormat:hVFL options:0 metrics:nil views:@{@"blueView":blueView}];
    [self.view addConstraints:hCons];
    
    //2.2垂直方向的约束
    NSString *vVFL = @"V:|-30-[blueView(50)]-30-[redView(==blueView)]";
    NSArray *vCons = [NSLayoutConstraint constraintsWithVisualFormat:vVFL options:NSLayoutFormatAlignAllRight metrics:nil views:@{@"blueView":blueView, @"redView":redView}];
    [self.view addConstraints:vCons];
    NSLayoutConstraint *redLeftCon = [NSLayoutConstraint constraintWithItem:redView attribute:NSLayoutAttributeLeft relatedBy:NSLayoutRelationEqual toItem:blueView attribute:NSLayoutAttributeCenterX multiplier:1.0 constant:0];
    [self.view addConstraint:redLeftCon];
}

@end
