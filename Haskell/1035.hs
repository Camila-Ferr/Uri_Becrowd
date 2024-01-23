leitorInt :: IO [Int]
leitorInt = fmap (map read.words) getLine

main :: IO ()
main = do
    valores <- leitorInt
    
    let a = valores !! 0
    let b = valores !! 1
    let c = valores !! 2
    let d = valores !! 3

    if (b > c) && (d > a) && (c+d > a+b) && (c > 0) && (d > 0) && (a`mod`2 == 0)
        then putStrLn "Valores aceitos"
    else
        putStrLn "Valores nao aceitos"
    


    