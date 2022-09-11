main :: IO ()
main = do
    fator1 <- getLine
    fator2 <- getLine

    let fator1Int = read fator1 :: Int
    let fator2Int = read fator2 :: Int

    putStrLn $ "SOMA = " ++show(fator1Int + fator2Int)