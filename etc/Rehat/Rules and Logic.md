## A motel should be managed by users
- One motel should be managed by at least one user

## Only privileged user account could manage a motel
- User should have `Admin` or `Superadmin` role to manage a motel
- Only `Admin` with permission could manage the motel
- `Superadmin` could manage any hotel, although it's not their main job

## Only vacant room could be reserved by users
- Room that is still within reservation period by other user could not be reserved until reservation has been checked out
- Room that is in maintenance could not be reserved until it the maintenance declared to be done

## Overstay would be charged extra
- Overstay indicated by checkout that hasn't been done before the end of reservation date. 
- Overstay charge amount is the same as per-night cost
- Overstay would be charged every 24 hours after reservation end date
- Motel room with overstaying user would be still tagged as `reserved` until the user do the checkout
