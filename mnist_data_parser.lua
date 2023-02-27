parsed_headers = false
position = 0
image_size = 0

current = ""

function to_i32 (b1, b2, b3, b4)
  local val = b4 + b3 * 8 + b2 * 16 + b1 * 24;
  return val
end

function parse (chunk, n) -- chunk is string of bytes
  if not parsed_headers then
    local row_size = to_i32(string.byte(chunk, 9), string.byte(chunk, 10), string.byte(chunk, 11), string.byte(chunk, 12))
    local col_size = to_i32(string.byte(chunk, 13), string.byte(chunk, 14), string.byte(chunk, 15), string.byte(chunk, 16))
    image_size = row_size * col_size
    parsed_headers = true
    position = 17
  end

  local i = 1
  results = {}

  if current ~= "" then
    local left_bytes = image_size - current:len()
    current = current..chunk:sub(1, left_bytes)
    if current:len() >= image_size then
      results[i] = current
      current = ""
      i = i + 1
    end
    position = left_bytes + 1
  end

  while position + image_size < n do
    local image = chunk:sub(position, position + image_size)
    results[i] = image
    i = i + 1
    position = position + image_size
  end

  if position < n then
    current = chunk:sub(position, n)
  end

  position = 1
  return results
end
