palindrone [] = error "empty list"
palindrone [x] = True
palindrone [x,y] = if x == y then True else False
palindrone (x:xs) = if (last xs) == x then palindrone (init xs) else False
