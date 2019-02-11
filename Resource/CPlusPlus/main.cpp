//
//  main.cpp
//  CPlusPlusDev
//
//  Created by ZhuHong on 2018/10/14.
//  Copyright © 2018年 CoderHG. All rights reserved.
//

#include <iostream>

using namespace std;

struct HGStruct {
    int m_id;
    int m_age;
    int m_height;
    
    virtual void display() {
        cout << "display" << this->m_id << endl;
        cout << "display" << this->m_age << endl;
    }
    
    
    void ttt() {
        cout << "ttt" << this << endl;
    }
};

struct HGGood : HGStruct {
    int m_name;
};

struct HGStruct01 {
public:
    static int mm_id;
public:
    int m_id;
    int m_age;
    int m_height;
    
    void display01() {
        cout << "display01" << this->m_id << endl;
        cout << "display01" << this->m_age << endl;
    }
};

int HGStruct01::mm_id = 230;


struct HGStruct02 : HGStruct01 {
    void test() {
        cout << "this.mm_id = " << this->mm_id << endl;
    }
    
};

struct HGObject {
    int m_a;
public:
    HGObject() {
        cout << "HGObject" << endl;
    }
};




int main(int argc, const char * argv[]) {
    HGObject obj;
//    HGStruct01::mm_id = 89;
//    HGStruct01::mm_id = 45;
    
//    HGStruct01 sss;
//    sss.mm_id = 89;
//
//    cout << "sss.mm_id = " << sss.mm_id << endl;
    
    HGStruct02 st;
    st.test();
    HGStruct02::mm_id = 90;
    st.test();
    
    HGGood good;
    good.m_id = 1;
    good.m_age = 2;
    good.m_height = 3;
    good.m_name = 4;
    
    cout << &good << ", " << &good.m_id << ", " << &good.m_age << ", " << &good.m_height << ", " << &good.m_name << endl;
    
    good.ttt();
    
    HGStruct s;
    s.m_id = 1;
    s.m_age = 2;
    s.m_height = 3;
    
    HGStruct01* pS = (HGStruct01*)&s;
    
    pS->display01();
    
    /**int a = 7;
    int b = 8;
    int c = a+b+4;
     */
    
    
    
    
    int a = 5;
    // int* p = &a;
    // *p = 6;
    int &rA = a;
    rA = 9;
    
    /*
    cout << a << endl;
    cout << *p << endl;
    cout << rA << endl;
    cout << sizeof(a) << endl;
    cout << sizeof(p) << endl;
    cout << sizeof(rA) << endl;
     */
    
    
    
    return 0;
}


/**
 CPlusPlusDev`main:
 0x100000f80 <+0>:  pushq  %rbp
 0x100000f81 <+1>:  movq   %rsp, %rbp
 0x100000f84 <+4>:  xorl   %eax, %eax
 0x100000f86 <+6>:  leaq   -0x14(%rbp), %rcx
 0x100000f8a <+10>: movl   $0x0, -0x4(%rbp)
 0x100000f91 <+17>: movl   %edi, -0x8(%rbp)
 0x100000f94 <+20>: movq   %rsi, -0x10(%rbp)
 0x100000f98 <+24>: movl   $0x5, -0x14(%rbp)
 ->  0x100000f9f <+31>: movq   %rcx, -0x20(%rbp)
 0x100000fa3 <+35>: movq   -0x20(%rbp), %rcx
 0x100000fa7 <+39>: movl   $0x6, (%rcx)
 0x100000fad <+45>: popq   %rbp
 0x100000fae <+46>: retq
 */
/**
 CPlusPlusDev`main:
 0x100000f80 <+0>:  pushq  %rbp
 0x100000f81 <+1>:  movq   %rsp, %rbp
 0x100000f84 <+4>:  xorl   %eax, %eax
 0x100000f86 <+6>:  movl   $0x0, -0x4(%rbp)
 0x100000f8d <+13>: movl   %edi, -0x8(%rbp)
 0x100000f90 <+16>: movq   %rsi, -0x10(%rbp)
 0x100000f94 <+20>: movl   $0x5, -0x14(%rbp)
 0x100000f9b <+27>: movl   $0x6, -0x18(%rbp)
 0x100000fa2 <+34>: movl   -0x14(%rbp), %edi
 0x100000fa5 <+37>: addl   -0x18(%rbp), %edi
 0x100000fa8 <+40>: addl   $0x3, %edi
 0x100000fab <+43>: movl   %edi, -0x1c(%rbp)
 ->  0x100000fae <+46>: popq   %rbp
 0x100000faf <+47>: retq
 */

/**
 CPlusPlusDev`main:
 0x100000f80 <+0>:  pushq  %rbp
 0x100000f81 <+1>:  movq   %rsp, %rbp
 0x100000f84 <+4>:  xorl   %eax, %eax
 0x100000f86 <+6>:  movl   $0x0, -0x4(%rbp)
 0x100000f8d <+13>: movl   %edi, -0x8(%rbp)
 0x100000f90 <+16>: movq   %rsi, -0x10(%rbp)
 0x100000f94 <+20>: movl   $0x7, -0x14(%rbp)
 0x100000f9b <+27>: movl   $0x8, -0x18(%rbp)
 ->  0x100000fa2 <+34>: movl   -0x14(%rbp), %edi
 0x100000fa5 <+37>: addl   -0x18(%rbp), %edi
 0x100000fa8 <+40>: addl   $0x4, %edi
 0x100000fab <+43>: movl   %edi, -0x1c(%rbp)
 0x100000fae <+46>: popq   %rbp
 0x100000faf <+47>: retq
 */
