<?php


/**
    二分查找
 * 又称折半查找，二分查找算法要求数据必须是有序的
 */
function BinarySearch($target, $arr){


    $length = count($arr);
    $start = 0;
    $end = $length - 1;

    while ($start <= $end){
        $mid= floor(($start + $end)/2);


        if($target == $arr[$mid]){
            return $arr[$mid];
        }else if($target < $arr[$mid]){
            $end =  $mid -1;
        }else{
            $start =  $mid + 1;
        }
    }

    return -1;
}


$arr = [1,2,3,4,5,8,9];
$target = 6;

$find = BinarySearch($target, $arr);
print_r($find);
