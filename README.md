# Central Coin

Did you ever wanted to feel like the **Feds** or any **central bank**? We're here for you.  
Print some currency out of thin air using **Central Coin**. A platform where anyone can create a **worthless currency**, invite friends and start trading it together.

## API Routes 

### Auth
All route should be protected by a JWT, the JWT must contain USER_ID, COIN_ID to be able to authorize or not some actions.

`POST /login`
- usernam/password combination to log in, using simple SQL table and encrypted pass.

`POST /signup`
- usernam/password combination to sign up, using simple SQL table and encrypted pass.

### Transactions
These routes are coin-related and allowed to authorized users only, COIN_ID and USER_ID should be present in JWT

`GET /transactions` 
- list all transactions on a given coin

`POST /transaction`
- Create a transaction on a given coin

`GET /transaction`
- Retrieve a transaction on a given coin

`SOCKET /TRANSACTION_EVENT/{coin_id}`
- Emit transaction event to users connected to socket.

### Users
These routes are coin-related, COIN_ID should be present in JWT

`GET /users` 
- list all users using a given coin (COIN_ID in JWT)

`GET /user/{USER_ID}`
- Retrieve a user using a given coin (COIN_ID in JWT)

### Coin
`POST /coin`
- Creates a new coin and assign USER as ADMIN to it

This route is allowed only by 
`GET /coin/invite`
- Retrieve invitation link to allow a user to join a COIN (COIN_ID, USER_ID om JWT) 
