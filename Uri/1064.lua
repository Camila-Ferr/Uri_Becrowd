local positivos = 0
local media = 0

for i = 1, 6 do
  local numero = tonumber(io.read())
  if numero > 0 then
    positivos = positivos + 1
    media = media + numero
  end
end

media = media / positivos

print(positivos .. " valores positivos")
print(string.format("%.1f", media))