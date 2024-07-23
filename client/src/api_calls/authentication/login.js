export async function login (email, password) {
  const res = await fetch('/api/account/login', {
    method: 'POST',
    body: JSON.stringify({
      email,
      password
    }),
    headers: {
      'Content-type': 'application/json; charset=UTF-8'
    }
  })
  console.log(res.json())
  return res
}
