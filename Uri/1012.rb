a, b, c = gets.split.map(&:to_f)
puts "TRIANGULO: #{'%.3f' % ((a*c)/2)}"
puts "CIRCULO: #{'%.3f' % (3.14159*c*c)}"
puts "TRAPEZIO: #{'%.3f' % (((a+b)*c)/2)}"
puts "QUADRADO: #{'%.3f' % (b*b)}"
puts "RETANGULO: #{'%.3f' % (a*b)}"