def fatorial(n)
    if n == 0 || n == 1
      return 1
    else
      return n * fatorial(n - 1)
    end
  end
  
  n = gets.to_i
  puts fatorial(n)