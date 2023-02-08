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
 