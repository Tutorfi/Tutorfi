export async function logout () {
  fetch('/api/account/logout', {
    method: 'GET',
    body: JSON.stringify({}),
    headers: {
      'Content-type': 'application/json; charset=UTF-8'
    }
  })
    .then((response) => response.json())
    .then((json) => { console.log(json) })
}
