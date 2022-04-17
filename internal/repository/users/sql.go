package users

const qryValidApiKey = `
	select exists (
		select true from users_api where api_key = $1
	);
`
