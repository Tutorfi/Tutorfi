import { useNavigate } from '@solidjs/router'
import { Show, createEffect, createSignal } from 'solid-js'

function loading(){
    return (
        <h1>Loading...</h1>
    );
}

function RouteGuard (props) {
    const navigate = useNavigate()
    console.log(props);

    const [auth, setAuth] = createSignal(false);
    fetch().then(/*Converting to json*/).then(/*manpulate the json*/).catch();
    
    return (
        <>
            <Show when={auth()} fallback={loading}>
                {props.children}
            </Show>
        </>);
}

export default RouteGuard
