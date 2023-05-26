def fatorial(n)
    if n == 0 || n == 1
      return 1
    else
      return n * fatorial(n - 1)
    end
  end
  
  loop do
    soma = 0
    input = gets
    break if input.nil?
  
    fat1, fat2 = input.chomp.split.map(&:to_i)
    soma += fatorial(fat1) + fatorial(fat2)
    puts soma
  end