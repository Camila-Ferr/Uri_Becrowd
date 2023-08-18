main :: IO ()
main = do
  entrada <- getLine
  let dias = read entrada :: Int
  putStrLn $ show(dias `div` 365) ++ " ano(s)"
  putStrLn $ show((dias `mod` 365) `div` 30) ++ " mes(es)"
  putStrLn $ show((dias `mod` 365) `mod` 30) ++ " dia(s)"