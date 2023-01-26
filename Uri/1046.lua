inicio, fim = io.read("*number", "*number")
if (fim == inicio) or (inicio > fim) then
  fim = 24+fim
end

print("O JOGO DUROU ".. math.abs(fim-inicio) .. " HORA(S)")