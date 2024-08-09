// create account
export async function register (fname, lname, email, username, password) {
  const res = await fetch('/api/account/create', {
    method: 'POST',
    body: JSON.stringify({
      fname,
      lname,
      email,
      username,
      password
    }),
    headers: {
      'Content-type': 'application/json; charset=UTF-8'
    }
  })
  return res
}
