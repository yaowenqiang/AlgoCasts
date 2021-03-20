<?php
function fib($n)
{
    static $count = 0;
    $count++;
    var_dump($count);

    if($n == 0 || $n == 1) {
        return $n;
    }
    return fib ($n-2) + fib($n-1);
}

var_dump('result:' . fib(10));
?>
