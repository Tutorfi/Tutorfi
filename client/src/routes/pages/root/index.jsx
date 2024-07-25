import { lazy, Switch, Match, createSignal } from 'solid-js'
import RouteGuard from '../../RouteGuard'
import { verifyCookie } from '../../../api_calls/authentication/verifyCookie'
const Landing = lazy(() => import('./Landing'))
const UserHome = lazy(() => import('../user/index'))

function root () {

  document.title = 'Tutorfi'
  const [auth, setAuth] = createSignal(false)
  // Change check whether the user is auth
  const verify = async () => { 
    const response = await verifyCookie()
      if (!response.ok) {
        setAuth(false)
      }else {
        setAuth(true)
      }
    }
    verify()
  return (
    <>
      <Switch>
        <Match when={!auth()}>
          <Landing />
        </Match>
        <Match when={auth()}>
          <UserHome />
        </Match>
      </Switch>
    </>
  )
}

export default root
