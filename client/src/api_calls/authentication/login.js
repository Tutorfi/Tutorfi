// verify
export async function login (email, password) {
  fetch('/api/account/verify', {
    method: 'POST',
    body: JSON.stringify({
      email,
      password
    }),
    headers: {
      'Content-type': 'application/json; charset=UTF-8'
    }
  })
    .then((response) => response.json())
    .then((json) => { console.log(json) })
}
