<?php

/**
冒泡排序
从小到大排序

时间复杂度： O(n^2)
 */

function BubbleSort(array $arr){
    $length = count($arr);
    if($length <= 1){
        return $arr;
    }

    for ($i = 0; $i < $length-1; $i ++){
        for ($j = $i + 1; $j < $length; $j ++){
            if($arr[$i] > $arr[$j]){
                $tmp = $arr[$i];
                $arr[$i] = $arr[$j];
                $arr[$j] = $tmp;
            }
        }
    }

    return $arr;
}


$arr = [8,2,1,5,3,7,9];

$res = BubbleSort($arr);

print_r($res);