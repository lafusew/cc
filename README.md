# Central Coin

## Routes 

All route should be protected by a JWT, the JWT must contain USER_ID, COIN_ID[] to be able to authorize or not some actions.

`POST /login`
- usernam/password combination to log in, using simple SQL table and encrypted pass.

`POST /signup`
- usernam/password combination to sign up, using simple SQL table and encrypted pass.

`GET /transactions/{COIN_ID}` 
- list all transactions on a given coin

`POST /transaction/{COIN_ID}`
- Create a transaction on a given coin