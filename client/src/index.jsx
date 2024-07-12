import { lazy } from 'solid-js'
import { render } from 'solid-js/web'
import { Router, Route } from '@solidjs/router'

// tailwind + flowbite
import './index.css'
import 'flowbite'

// lazy load pages
const RootPage = lazy(() => import('./routes/pages/root/index'))
const OrgsPage = lazy(() => import('./routes/pages/orgs/index'))
const GroupsPage = lazy(() => import('./routes/pages/groups/index'))
const NotFound = lazy(() => import('./routes/pages/NotFound'))
const CalendarPage = lazy(() => import('./components/calendar/calendar'))

const root = document.getElementById('root')

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    'Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got misspelled?'
  )
}

// router
render(() => (
  <Router>
    <Route path="/" component={RootPage} />
    <Route path="/dashboard" component={RootPage} />
    <Route path="/orgs" component={OrgsPage} />
    <Route path="/groups" component={GroupsPage} />
    <Route path="/calendar" component={CalendarPage} />
    <Route path="*404" component={NotFound} />
  </Router>
), document.getElementById('root'))
