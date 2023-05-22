n = gets.to_i

for i in 1..n do
  soma = 0
  texto = gets.chomp
  quantidade_letras = texto.length

  if quantidade_letras > 3
    puts '3'
  else
    soma += 1 if texto[0].ord == 116
    soma += 1 if texto[1].ord == 119
    soma += 1 if texto[2].ord == 111

    if soma >= 2
      puts '2'
    else
      puts '1'
    end
  end
end