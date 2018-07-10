package model

const selectUsersQuery = `SELECT user_id, full_name, msisdn, user_email, birth_date, create_time, update_time
FROM public.ws_user
WHERE birth_date is not null AND update_time is not null
LIMIT 25;`
