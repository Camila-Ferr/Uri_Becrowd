total = 0
for i in 1..2 do
  cod, quantidade, preco = gets.split.map(&:to_f)
  total += quantidade * preco
end
puts "VALOR A PAGAR: R$ #{'%.2f' % total}"