import Control.Monad
import Text.Printf

calcula :: Int -> [Int] -> Int -> IO() 
calcula valor nota  index = do

    let quantidade = valor `div` nota!!index
    
    putStr $ printf "%d nota(s) de R$ %d,00\n" quantidade (nota !! index)

    when (index < 6) $ do 
        calcula (valor - (quantidade * nota !! index)) nota (index+1)
            

main :: IO ()
main = do

    valor <- getLine

    let valorI = (read valor :: Int)
    putStr $ printf "%d\n" valorI

    calcula valorI [100,50,20,10,5,2,1] 0 
    