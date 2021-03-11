let chunk x y = if x == [] then [] else if  length x < y then [x] else (take y x) : (chunk (drop y x) y)
