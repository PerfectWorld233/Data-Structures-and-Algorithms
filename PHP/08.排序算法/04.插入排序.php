<?php

/**
 插入排序（Insertion Sort）
    是一种简单直观的排序算法。
    它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，
    找到相应位置并插入。插入排序在实现上，通常采用in-place排序（即只需用到O(1)的额外空间的排序），
    因而在从后向前扫描过程中，需要反复把已排序元素逐步向后挪位，为最新元素提供插入空间。

    时间复杂度： O(N^(1-2))
 *
 *   可以理解为打字牌，通过插入依次调整牌的大小，使之有序
 */



function InsertSort(array $arr){

    $len = count($arr);
    for ($i = 0; $i < $len; $i ++){
        $key = $arr[$i];    // 当前值
        $pos = $i;          // 当前位置

        while($pos > 0 && $arr[$pos - 1] > $key){
            $arr[$pos] = $arr[$pos - 1];
            $pos = $pos - 1;
        }
        $arr[$pos] = $key;      // 找到单前值的位置

    }

    return $arr;
}


$arr = [2,4,1,5,9,7,8];

$res = InsertSort($arr);
print_r($res);