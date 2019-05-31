//
//  main.cpp
//  HGPlusRandom
//
//  Created  by hong.zhu on 2019/5/31.
//  Copyright © 2019 INC. All rights reserved.
//

#include <iostream>
#include "CryptoUtils.hpp"

int main(int argc, const char * argv[]) {
    CryptoUtils *cUtils = new CryptoUtils();
    
    for (int i=0; i<5; i++) {
        std::cout << cUtils->get_uint32_t();
        std::cout << "\n";
    }
    std::cout << "=======get_uint32_t=======\n";
    
    for (int i=0; i<5; i++) {
        std::cout << cUtils->get_uint64_t();
        std::cout << "\n";
    }
    std::cout << "========get_uint64_t======\n";
    
    for (int i=0; i<5; i++) {
        std::cout << cUtils->get_uint8_t();
        std::cout << "\n";
    }
    std::cout << "========get_uint8_t======\n";
    
    for (int i=0; i<5; i++) {
        std::cout << cUtils->get_uint16_t();
        std::cout << "\n";
    }
    std::cout << "========get_uint16_t======\n";
    
    
    return 0;
}
