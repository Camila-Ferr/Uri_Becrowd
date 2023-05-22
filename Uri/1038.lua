local produtos = { 4.00, 4.50, 5.00, 2.00, 1.50 }
local entrada = io.read("*line")
local indice, valor = string.match(entrada, "(%S+)%s+(%S+)")
indice = tonumber(indice)
valor = tonumber(valor)

print(string.format("Total: R$ %.2f", produtos[indice] * valor))
