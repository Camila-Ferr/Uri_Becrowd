import Text.Printf

main :: IO ()
main = do

    nota1 <- getLine
    nota2 <- getLine
    
    let double1 = (read nota1 :: Double)
    let double2 = (read nota2 :: Double)

    let media = (3.5 * double1 + 7.5 * double2)/11

    putStrLn $ "MEDIA = " ++ printf "%.5f" media