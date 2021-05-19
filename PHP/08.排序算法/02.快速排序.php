<?php


/**
快速排序

1、选定一个基准值(任意选,以第一个为例)
2、定义左右指针指向左右两端
3、左指针往右移动，
如果遇到大于基准值的数就把它和右指针的值调换位置,
然后左指针不动,右指针开始向左移动,
如果遇到小于基准值的数就把他和左指针的值调换位置,然后开始移动左指针,以此类推,知道左右指针相遇
4、递归上述过程直到排序结束
 */


function QuickSort(array $arr){

    $length = count($arr);
    if($length <= 1){
       return $arr;
    }


    $middle = $arr[0];

    $left = array();        // 接收小于中间值
    $right = array();       // 接收 大于中间值

    for ($i = 1; $i < $length; $i ++){
        if($i < $middle){
            $left[] = $arr[$i];
        }else{
            $right[] = $arr[$i];
        }
    }

    $left = QuickSort($left);
    $right = QuickSort($right);


    $res = array_merge($left, array($middle), $right);


    return $res;
}



$arr = [2,1,9,8,4,6,5];
$res = QuickSort($arr);
print_r($res);
