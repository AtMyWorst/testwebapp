parsed_headers = false
position = 0
image_size = 0

current = ""

function to_i32 (b1, b2, b3, b4)
  local val = b4 + b3 * 8 + b2 * 16 + b1 * 24;
  return val
end

function parse (chunk, n) -- chunk is string o