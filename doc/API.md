### 1. LOGIN:
```sh
curl --location 'localhost:8080/api/v1/user/login' \
--header 'Content-Type: application/json' \
--data '{
    "phone": "+99892720842",
    "otp": 1233
}'
```
#### Response OBJ:
```sh
{
    "message": "Login successful",
    "status": true,
    "userInfo": {
        "image": null,
        "name": "Test",
        "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZSI6Iis5OTg5MjcyMDg0MiIsImV4cCI6MTcyODM3OTY1NH0.4n1NdakchpAfQJessw0KYYqKfSDtYGDO7FiW7x4l14Y",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZSI6Iis5OTg5MjcyMDg0MiIsImV4cCI6MTczNTU1MDg1NH0.RSB_3EMsY9AIIfrlLrGY2KuLrFNTcguBLGIVnHcmAAE"
    }
}
```

### 2. REGISTER:
```sh
curl --location 'localhost:8080/api/v1/user/register' \
--header 'Content-Type: application/json' \
--data '{
    "phone": "+99892720842",
    "name": "Test",
    "otp": 1234,
    "userType": 4
}'
```

### 3. SEND-OTP:
```sh
curl --location 'localhost:8080/api/v1/user/send-otp' \
--header 'Content-Type: application/json' \
--data '{
    "phone": "+99892720841"
}'
```

### 4. GET-ADDRESS:
```sh
curl --location 'localhost:8080/api/v1/address/get-address' \
--header 'Content-Type: application/json' \
--data '{
    "pincode": 400001
}'
```
> OR
```sh
curl --location 'localhost:8080/api/v1/address/get-address' \
--header 'Content-Type: application/json' \
--data '{
    "state": "maharashtra"
}'
```

#### Response OBJ:
```sh
{
    "data": [
        {
            "id": 9659,
            "pincode": 400001,
            "city": "Mumbai",
            "district": "Mumbai",
            "state": "Maharashtra",
            "country_id": 101,
            "country": {
                "id": 0,
                "sortname": "",
                "name": ""
            }
        },{},{},...,
	]
}
```
### 4. UPDATE-ADDRESS:
```sh
curl --location 'localhost:8080/api/v1/user/update-address/1' \
--header 'Content-Type: application/json' \
--data '{
    "house": "C/301",
    "area": "New link road, Evershine",
    "landmark": "D-Mart",
    "sd_address_id": 123
}'
```