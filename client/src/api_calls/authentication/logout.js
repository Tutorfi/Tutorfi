export async function logout () {
  return await fetch('/api/account/logout', {
    method: 'GET'
  })
}
