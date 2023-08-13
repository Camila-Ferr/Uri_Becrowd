def primo?(numero)
    return false if numero <= 1
  
    (2..Math.sqrt(numero)).each do |i|
      return false if numero % i == 0
    end
  
    true
  end
  
  def separar_inteiro(numero)
    digitos = numero.to_s.chars.map(&:to_i)
    digitos.each do |digito|
      return false unless primo?(digito)
    end
    true
  end
  
  loop do
    entrada = gets
    break if entrada.nil? || entrada.chomp.empty?
  
    n = entrada.to_i
    saida = 'Nada'
    if primo? n
      saida = 'Primo'
      saida = 'Super' if separar_inteiro(n)
    end
    puts saida
  end  