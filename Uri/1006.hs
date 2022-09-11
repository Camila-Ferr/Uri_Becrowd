import Text.Printf

main :: IO ()
main = do

    nota1 <- getLine
    nota2 <- getLine
    nota3 <- getLine

    let double1 = (read nota1 :: Double)
    let double2 = (read nota2 :: Double)
    let double3 = (read nota3 :: Double)


    let media = (2 * double1 + 3 * double2 + 5 * double3)/10

    putStrLn $ "MEDIA = " ++ printf "%.1f" media