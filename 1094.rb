coelho,rato,sapo,total=0,0,0,0
n = gets.to_i
for i in 1..n do
  texto = gets.chomp
  inteiro, char = texto.split(' ', 2)
  inteiro = inteiro.to_i

  coelho+=inteiro if char[0].ord == 67
  rato+=inteiro if char[0].ord == 82
  sapo+=inteiro if char[0].ord == 83
  total+=inteiro
end
puts "Total: #{total} cobaias"
puts "Total de coelhos: #{coelho}"
puts "Total de ratos: #{rato}"
puts "Total de sapos: #{sapo}"
puts "Percentual de coelhos: #{'%.2f' % ((coelho.to_f / total) * 100)} %"
puts "Percentual de ratos: #{'%.2f' % ((rato.to_f / total) * 100)} %"
puts "Percentual de sapos: #{'%.2f' % ((sapo.to_f / total) * 100)} %"