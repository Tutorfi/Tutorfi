import { useNavigate } from '@solidjs/router'
import { Show, createSignal } from 'solid-js'
import { verifyCookie } from '../api_calls/authentication/verifyCookie'

function loading(){
    return (
        <h1>Loading...</h1>
    );
}

function RouteGuard (props) {
  const navigate = useNavigate()
  const [auth, setAuth] = createSignal(false)

  const verify = async () => { 
    const response = await verifyCookie()
    if ( !response.ok ) {
      navigate('/', { replace: true })
      setAuth(false)
    }else {
      setAuth(true)
    }
  }
  verify()

  return (
  <>
    <Show when={auth()} fallback={loading}>
      {props.children}
    </Show>
  </>);
}

export default RouteGuard
