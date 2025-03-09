--下面的参数分别以key为ip和userId为前缀以区分
--key: The key to limit
--count:需要获取的token数量
--capacity:令牌漏桶的容量
--rate:令牌生成的速度，每秒生成多少个令牌
--leave_tokens:桶中上次剩余的令牌数量
--last_time:上一次获取令牌的时间，单位毫秒
--返回:1为获取成功，-1为获取失败

local function getTokens(key, count)

    local bucket_info = redis.pcall("HMGET", key, "last_time", "capacity", "rate", "current_tokens")
    local last_time = tonumber(bucket_info[1])

    local capacity = tonumber(bucket_info[2])

    local rate = tonumber(bucket_info[3])

    local leave_tokens = tonumber(bucket_info[4])
    --本次计算时的令牌数量
    local local_tokens = 0

    local time = redis.call("TIME")
    local now_time = time[1] * 1000000 + time[2]
    now_time = now_time / 1000

    if last_time ~= nil then
        --从上次获取令牌到现在生成的令牌数量
        local generated_tokens = math.floor(((now_time- last_time) / 1000) * rate)
        local_tokens = math.min(leave_tokens + generated_tokens, capacity)
    else
        --第一次获取令牌
        redis.pcall("HSET", key, "last_time", now_time)
        local_tokens = capacity

    end

    local result = -1

    if local_tokens >= tonumber(count) then
        --令牌充足
        redis.pcall("HSET", key, "current_tokens", local_tokens - count)
        redis.pcall("HSET", key, "last_time", now_time)
        result = 1
    else
        --令牌不足
        result = -1
    end
    return result

end

local function init(key, capacity, rate)
    local bucket_info = redis.pcall("HMGET", key, "last_time", "capacity", "rate", "current_tokens")
    local org_capacity = tonumber(bucket_info[2])
    local org_rate = tonumber(bucket_info[3])

    if org_capacity == nil or org_rate ~= rate or org_capacity ~= capacity then
        redis.pcall("HMSET", key, "capacity", capacity, "rate", rate, "current_tokens", capacity, "capacity", capacity)
    end
    --每次访问都设置过期时间
    redis.pcall("EXPIRE", key, 30)

    return 1

end

local key = KEYS[1]
local method_name = ARGV[1]
if method_name == "getTokens" then
    return getTokens(key, ARGV[2])
elseif method_name == "init" then
    return init(key, tonumber(ARGV[2]), tonumber(ARGV[3]))
else
    return -1
end




