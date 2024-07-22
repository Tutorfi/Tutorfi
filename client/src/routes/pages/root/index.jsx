import { lazy, Switch, Match } from 'solid-js'
import RouteGuard from '../../RouteGuard'
const Landing = lazy(() => import('./Landing'))
const UserHome = lazy(() => import('../user/index'))

function root () {

  document.title = 'Tutorfi'

  const token = sessionStorage.getItem('token')

  return (
    <>
      <Switch>
        <Match when={!token}>
          <Landing />
        </Match>
        <Match when={token}>
          <UserHome />
        </Match>
      </Switch>
    </>
  )
}

export default root
