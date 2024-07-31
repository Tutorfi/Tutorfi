// verifies the user login credentials
export async function verifyCookie () {
  const response = await fetch('/api/account/verify', {
    method: 'GET'
  })
  console.log(response.json())
  return response
}
