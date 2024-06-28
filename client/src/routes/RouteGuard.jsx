import { useNavigate } from '@solidjs/router'
import { createEffect } from 'solid-js'

function RouteGuard () {
  const navigate = useNavigate()
  const token = sessionStorage.getItem('token')

  createEffect(() => {
    if (!token) {
      navigate('/', { replace: true })
    }
  })
}

export default RouteGuard
