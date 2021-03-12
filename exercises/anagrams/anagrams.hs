quicksort :: Ord a => [a] -> [a]
quicksort []     = []
quicksort (p:xs) = (quicksort lesser) ++ [p] ++ (quicksort greater)
    where
        lesser  = filter (< p) xs
        greater = filter (>= p) xs

quicksort' [] = []
quicksort' (p:xs) = (quicksort (filter (< p) xs)) ++ [p] ++ (quicksort (filter (>= p) xs))

anagram x y = (quicksort x) == (quicksort y)

