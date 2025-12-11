local http = require "resty.http"
local cjson = require "cjson"

local TokenValidatorHandler = {}

TokenValidatorHandler.PRIORITY = 1000
TokenValidatorHandler.VERSION = "1.0.0"

local function get_bearer_token(header)
  if not header then
    return nil
  end

  local token = header:match("^[Bb]earer%s+(.+)$")
  return token
end

function TokenValidatorHandler:access(conf)
  if kong.request.get_method() == "OPTIONS" then
    return
  end

  local token = get_bearer_token(kong.request.get_header("Authorization"))

  if not token then
    return kong.response.exit(401, {
      message = "No Bearer token found in Authorization header"
    }, {
      ["Content-Type"] = "application/json"
    })
  end

  local auth_service_url = os.getenv("AUTHENTICATION_SERVICE_URL") or "http://host.docker.internal:8081"
  local userinfo_url = auth_service_url .. "/api/authentication/v1/authentication/userinfo"
  local httpc = http.new()
  httpc:set_timeout(2000)

  local res, err = httpc:request_uri(userinfo_url, {
    method = "GET",
    headers = {
      ["Authorization"] = "Bearer " .. token,
    },
  })

  if not res then
    kong.log.err("Failed to verify token: ", err)
    return kong.response.exit(500, {
      message = "Token verification service unavailable"
    }, {
      ["Content-Type"] = "application/json"
    })
  end

  if res.status ~= 200 then
    return kong.response.exit(401, {
      message = "Invalid or expired token"
    }, {
      ["Content-Type"] = "application/json"
    })
  end

  local function base64url_decode(input)
    input = input:gsub("-", "+"):gsub("_", "/")
    local remainder = #input % 4
    if remainder > 0 then
      input = input .. string.rep("=", 4 - remainder)
    end
    return ngx.decode_base64(input)
  end

  local parts = {}
  for part in token:gmatch("[^%.]+") do
    table.insert(parts, part)
  end

  if #parts >= 2 then
    local payload_json = base64url_decode(parts[2])
    if payload_json then
      local ok, payload = pcall(cjson.decode, payload_json)
      if ok and payload then
        if payload.sub then
          kong.service.request.set_header("X-User-Sub", payload.sub)
        end

        if payload.email then
          kong.service.request.set_header("X-User-Email", payload.email)
        end
      end
    end
  end

  return
end

return TokenValidatorHandler

