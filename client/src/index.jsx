import { lazy } from 'solid-js'
import { render } from 'solid-js/web'
import { HashRouter, Route } from '@solidjs/router'

const Landing = lazy(() => import('./routes/pages/Landing'))
const NotFound = lazy(() => import('./routes/pages/NotFound'))

const root = document.getElementById('root')

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    'Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got misspelled?'
  )
}

render(() => (
  <HashRouter>
    <Route path="/" component={Landing} />
    <Route path="*404" component={NotFound} />
  </HashRouter>
), document.getElementById('root'))
