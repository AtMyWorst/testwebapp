
parsed_headers = false
position = 0

function parse (chunk, n) -- chunk is string of bytes
  if not parsed_headers then
    parsed_headers = true
    position = 9
  end

  local i = 1
  results = {}

  while position < chunk:len() do
    results[i] = chunk:sub(position, position)
    i = i + 1
    position = position + 1
  end

  position = 1
  return results
end