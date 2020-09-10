

# non existent endpoint
curl "http://localhost:8090/dne"

# example api
curl "http://localhost:8090/api/example/hello" | json_pp
curl "http://localhost:8090/api/example/get_headers" | json_pp
curl "http://localhost:8090/api/example/get_get_params?a=1&b=2" | json_pp

# user api
curl "http://localhost:8090/api/user/get_user" | json_pp

# math api
curl "http://localhost:8090/api/math/add?n1=1&n2=2" | json_pp