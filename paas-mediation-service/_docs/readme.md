
# How to generate swagger doc manually

1. install swag/cmd v1.8.12 or above
   ```
   go install github.com/swaggo/swag/cmd/swag@v1.8.12
   ```
2. from withing paas-mediation-service folder execute swag init to generate swagger.json file
   ```
   swag init 
   ```
3. install https://github.com/go-swagger/go-swagger
4. generate MD doc from swagger.json file
   ```
   swagger generate markdown -f ./docs/swagger.json --output ./../docs/rest_api.md
 
   ```
