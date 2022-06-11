# Central Coin

## API Routes 

### Auth

All route should be protected by a JWT, the JWT must contain USER_ID, COIN_ID to be able to authorize or not some actions.

`POST /login`
- usernam/password combination to log in, using simple SQL table and encrypted pass.

`POST /signup`
- usernam/password combination to sign up, using simple SQL table and encrypted pass.

### Transactions

`GET /transactions` 
- list all transactions on a given coin (COIN_ID in JWT)

`POST /transaction`
- Create a transaction on a given coin (COIN_ID in JWT)

`GET /transaction`
- Retrieve a transaction on a given coin (COIN_ID in JWT)

`SOCKET /TRANSACTION_EVENT/{coin_id}`
- Emit transaction event to users connected to socket.

### Users

`GET /users` 
- list all users using a given coin (COIN_ID in JWT)

`GET /user/{USER_ID}`
- Retrieve a user using a given coin (COIN_ID in JWT)

### Coin

`POST /coin`
- Creates a new coin and assign USER as ADMIN to it

`GET /coin/invite`
- Retrieve invitation link to allow a user to join a COIN (COIN_ID, USER_ID om JWT) 