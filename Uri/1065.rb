soma = 0
for i in 1..5 do
  n = gets.to_i
  if (n%2==0) then 
    soma +=1
  end
end
puts "#{soma} valores pares"