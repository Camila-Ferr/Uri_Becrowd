loop do
    input = gets.chomp
    break if input == '0'
  
    quantidade = input.length
    soma = 0
    fatorial = 1
  
    for i in 1..quantidade do
      fatorial *= i
      soma += input[quantidade - i].to_i * fatorial
  
    end
  
    puts soma
  end