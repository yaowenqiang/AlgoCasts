reverseList' xs  = foldl (\x y -> y:x) [] xs

reverseList  [] = []
reverseList  xs = last xs : reverseList (init xs)
