import Text.Printf

leitorDouble :: IO [Double]
leitorDouble = fmap (map read.words) getLine

calculaR :: Double -> Double -> Double -> IO() 
calculaR a b delta = do

    let raiz1 = ((b * (-1)) + sqrt (delta))/ (2*a)
    let raiz2 = ((b * (-1)) - sqrt (delta))/ (2*a)
    
    putStrLn $ "R1 = " ++ printf "%.5f" raiz1
    putStrLn $ "R2 = " ++ printf "%.5f" raiz2
                    
    



main :: IO ()
main = do
    valores <- leitorDouble 
    let a = valores !! 0
    let b = valores !! 1
    let c = valores !! 2
    
    let delta = b**2 - 4*a*c
    
    if (delta < 0 || a == 0) then do 
        putStrLn $ printf "Impossivel calcular"
    else
        calculaR a b delta
    