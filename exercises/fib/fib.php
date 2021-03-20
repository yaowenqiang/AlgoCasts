<?php
function fib($n)
{
    static $count = 0;
    $count++;
    //var_dump($count);

    if($n == 0 || $n == 1) {
        return $n;
    }
    return memo('fib', $n-2) + memo('fib', $n-1);
}

/*
function memo($n)
{
    if($n == 0 || $n == 1) {
        return $n;
    }
    static $cache = [];
    if (isset($cache[$n])) {
        return memo[$n];
    } else {
        $result = memo($n);
        $cache[$n] = $result;
        return $result;
    }

}
 */

function fastFib($n)
{
    static $cache = [0, 1];
    static $count = 0;
    $count++;
    var_dump($count);
    static $self = __FUNCTION__;
    if (isset($cache[$n])) {
        return $cache[$n];
    } else {
        $result = $self($n-1) + $self($n - 2);
        //$result = "{__FUNCTION__}"($n-1) + "{__FUNCTION__}"($n - 2);
        $cache[$n] = $result;
        return $result;
    }
}

function slowFib($n)
{
    static $count = 0;
    $count++;
    var_dump($count);

    if($n == 0 || $n == 1) {
        return $n;
    }
    return slowFib($n-2) + slowFib($n-1);
}

//var_dump('result:' . fib(12));
var_dump('result:' . fastFib(10));
var_dump('result:' . slowFib(10));
?>
