export async function login (email, password, remember) {
  const res = await fetch('/api/account/login', {
    method: 'POST',
    body: JSON.stringify({
      email,
      password,
      remember,
    }),
    headers: {
      'Content-type': 'application/json; charset=UTF-8'
    }
  })
  console.log(res.json())
  return res
}
