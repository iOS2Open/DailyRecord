//
//  CryptoUtils.cpp
//  HGPlusRandom
//
//  Created  by hong.zhu on 2019/5/31.
//  Copyright © 2019 INC. All rights reserved.
//

#include "CryptoUtils.hpp"

CryptoUtils::CryptoUtils() {
}

uint32_t CryptoUtils::scramble32(uint32_t in,std::map<uint32_t/*IDX*/,uint32_t/*VAL*/>& VMap) {
    if(VMap.find(in)==VMap.end()){
        uint32_t V=get_uint32_t();
        VMap[in]=V;
        return V;
    }
    else{
        return VMap[in];
    }
}

CryptoUtils::~CryptoUtils() {
    if(eng!=nullptr){
        delete eng;
    }
}

void CryptoUtils::prng_seed(){
    using namespace std::chrono;
    std::uint_fast64_t ms = duration_cast< milliseconds >(system_clock::now().time_since_epoch()).count();
    // errs()<<format("std::mt19937_64 seeded with current timestamp: %" PRIu64"",ms)<<"\n";
    eng=new std::mt19937_64(ms);
}

void CryptoUtils::prng_seed(std::uint_fast64_t seed){
    // errs()<<format("std::mt19937_64 seeded with: %" PRIu64"",seed)<<"\n";
    eng=new std::mt19937_64(seed);
}

std::uint_fast64_t CryptoUtils::get_raw(){
    if(eng==nullptr){
        prng_seed();
    }
    return (*eng)();
}

uint32_t CryptoUtils::get_range(uint32_t min,uint32_t max) {
    if(max==0){
        return 0;
    }
    std::uniform_int_distribution<uint32_t> dis(min, max-1);
    return dis(*eng);
}
